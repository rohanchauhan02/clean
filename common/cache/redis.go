package cache

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis"
)

type redisCache struct {
	redisSess *redistrace.Client
}

func NewRedisCache(redisSess *redistrace.Client) redisCache {
	return redisCache{
		redisSess: redisSess,
	}
}

func (r redisCache) SetRedisValue(key *string, payload *string, ttl time.Duration) {
	r.redisSess.Set(*key, *payload, ttl)
}

func (r redisCache) GetRedisValue(key *string) *string {
	val, err := r.redisSess.Get(*key).Result()
	if err != nil {
		return nil
	}
	return aws.String(val)
}

func (r redisCache) RemoveRedisValue(key *string) *int64 {
	val, err := r.redisSess.Del(*key).Result()
	if err != nil {
		return nil
	}
	return aws.Int64(val)
}
