package limiter

import (
	config "AeromindGO/config/env"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisBucketService interface {
	GetClientKey(identifier string) string
	getBucketStatus(identifier string) 
}

type RedisBucketServiceImpl struct {
	redisClient *redis.Client
	defaultCapacity int
	defaultRefill int
	defaultTimeout int
	redisTimeout time.Duration
}

type BucketStatus struct{
	exists bool `json:"exists"`
}

func (bucket *RedisBucketServiceImpl) GetClientKey(identifier string) string{

	key := fmt.Sprintf("r1:tb:%s",identifier)

	return key

}

func (bucket *RedisBucketServiceImpl) getBucketStatus(identifier string){


	key:=bucket.GetClientKey(identifier)

	


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