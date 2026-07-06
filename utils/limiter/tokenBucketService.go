package limiter

import (
	config "AeromindGO/config/env"
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisBucketService interface {
	GetClientKey(identifier string) string
	ResetBucket(identifier string) error
	GetBucketStatus(identifier string) (*BucketStatus,error)
	IsRequestAllowed(identifier string, options *BucketOptions) (*BucketResponse, error)
	ListAllBuckets() ([]string, error)
}

type RedisBucketServiceImpl struct {
	redisClient *redis.Client
	defaultCapacity int
	defaultRefill int
	defaultTimeout int
	redisTimeout time.Duration
}

type BucketStatus struct{
	Exists          bool   `json:"exists"`
	Identifier      string `json:"identifier"`
	StoredTokens    int    `json:"storedTokens"`
	ProjectedTokens int    `json:"projectedTokens"`
	Capacity        int    `json:"capacity"`
	LastRefillTime  int64  `json:"lastRefillTime"`
	TTLMs           int64  `json:"ttlMs"`
	RedisKey        string `json:"redisKey"`
	Error           string `json:"error,omitempty"`
}

type BucketOptions struct {
    Capacity int
    Refill   int
    Timeout  int
    Cost     int
}

type BucketResponse struct {
    Allowed      bool
    Remaining    int
    Capacity     int
    RetryAfterMs int64
    Identifier   string
}

func (bucket *RedisBucketServiceImpl) GetClientKey(identifier string) string{

	key := fmt.Sprintf("r1:tb:%s",identifier)

	return key

}

func (bucket *RedisBucketServiceImpl) ResetBucket(identifier string) error {
	ctx := context.Background()
	key := bucket.GetClientKey(identifier)
	return bucket.redisClient.Del(ctx, key).Err()
}

func (bucket *RedisBucketServiceImpl) IsRequestAllowed(
	identifier string,
	options *BucketOptions,
) (*BucketResponse, error) {

	key := bucket.GetClientKey(identifier)
	ctx := context.Background()
	capacity := bucket.defaultCapacity
	refill := bucket.defaultRefill
	timeout := bucket.defaultTimeout
	cost := 1

	if options != nil {
		if options.Capacity != 0 {
			capacity = options.Capacity
		}

		if options.Refill != 0 {
			refill = options.Refill
		}

		if options.Timeout != 0 {
			timeout = options.Timeout
		}

		if options.Cost != 0 {
			cost = options.Cost
		}
	}

	refillRate := float64(refill) / float64(timeout)
	script := Lua_Scripter()

	rawResult, err := bucket.redisClient.Eval(
		ctx,
		script,
		[]string{key},
		capacity,
		refillRate,
		time.Now().UnixMilli(),
		cost,
	).Result()
	if err != nil {
		return nil, err
	}

	result, ok := rawResult.([]interface{})
	if !ok || len(result) < 3 {
		return nil, fmt.Errorf("unexpected rate limiter response: %v", rawResult)
	}

	allowed, err := strconv.Atoi(fmt.Sprint(result[0]))
	if err != nil {
		return nil, err
	}

	remaining, err := strconv.Atoi(fmt.Sprint(result[1]))
	if err != nil {
		return nil, err
	}

	retryAfterMs, err := strconv.ParseInt(fmt.Sprint(result[2]), 10, 64)
	if err != nil {
		return nil, err
	}

	return &BucketResponse{
		Allowed:      allowed == 1,
		Remaining:    remaining,
		Capacity:     capacity,
		RetryAfterMs: retryAfterMs,
		Identifier:   identifier,
	}, nil
}

func (bucket *RedisBucketServiceImpl) ListAllBuckets() ([]string, error) {

    var (
        cursor uint64
        keys []string
    )

    for {
        result, next, err := bucket.redisClient.Scan(
            context.Background(),
            cursor,
            "r1:tb:*",
            100,
        ).Result()

        if err != nil {
            return nil, err
        }

        keys = append(keys, result...)
        cursor = next

        if cursor == 0 {
            break
        }
    }

    return keys, nil
}

func (bucket *RedisBucketServiceImpl) GetBucketStatus(identifier string) (*BucketStatus,error){

	ctx := context.Background()
	key := bucket.GetClientKey(identifier)


	data, err := bucket.redisClient.HMGet(ctx, key, "tokens",
		"lastRefillTime").Result()


	if err != nil{
		return &BucketStatus{
			Exists: false,
			Identifier: identifier,
			Error: err.Error(),
		},nil
	}

	ttl,TTLerr:=bucket.redisClient.PTTL(ctx,key).Result()

	if TTLerr != nil{
		return nil,TTLerr
	}

	if data[0] == nil {
		return &BucketStatus{
			Exists: false,
			Identifier: identifier,
		},nil
	}

	tokens,_:=strconv.ParseFloat(data[0].(string),64)
	lastFill,_:=strconv.ParseInt(data[1].(string),10,64)

	refillRate:= float64(bucket.defaultRefill)/float64(bucket.defaultTimeout)
	elapsed:= time.Now().UnixMilli() - lastFill


	projected:=math.Min(
		float64(bucket.defaultCapacity),
		tokens+float64(elapsed)*refillRate,
	)
	return &BucketStatus{
		Exists        :  true,
		Identifier      :identifier,
		StoredTokens    :int(tokens),
		ProjectedTokens :int(projected),
		Capacity        :bucket.defaultCapacity,
		LastRefillTime  :lastFill,
		TTLMs           :ttl.Milliseconds(),
		RedisKey        :key,
	},nil

}

func NewRedisBucketService(client *redis.Client)RedisBucketService{
	return &RedisBucketServiceImpl{
		redisClient: client,
		defaultCapacity: config.GetInt("BUCKET_CAPACITY",100),
		defaultRefill: config.GetInt("BUCKET_REFILL",10),
		defaultTimeout: config.GetInt("BUCKET_TIMEOUT",60000),
		redisTimeout: time.Duration(config.GetInt("REDIS_TIMEOUT",60000)),
	}
}