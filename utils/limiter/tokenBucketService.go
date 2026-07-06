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
	resetBucket(identifier string) error
	getBucketStatus(identifier string) (*BucketStatus,error)
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

func (bucket *RedisBucketServiceImpl) GetClientKey(identifier string) string{

	key := fmt.Sprintf("r1:tb:%s",identifier)

	return key

}

func(bucket *RedisBucketServiceImpl) resetBucket(identifier string) error {

	ctx:=context.Background()
	key:=bucket.GetClientKey(identifier)
	return bucket.redisClient.Del(ctx,key).Err()

}

func (bucket *RedisBucketServiceImpl) getBucketStatus(identifier string) (*BucketStatus,error){

	ctx:=context.Background()
	key:=bucket.GetClientKey(identifier)


	data,err:=bucket.redisClient.HMGet(ctx,key,"tokens",
		"lastRefillTime",).Result()


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