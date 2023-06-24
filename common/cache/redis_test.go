package cache

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis"
)

var redisServer *miniredis.Miniredis
var redisClient *redistrace.Client
var redisSvc redisCache

func setup() {
	redisServer = mockRedis()
	redisClient = redistrace.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})

	redisSvc = redisCache{
		redisSess: redisClient,
	}
}
func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	return s
}
func teardown() {
	redisServer.Close()
}

func TestNewRedisCache(t *testing.T) {
	setup()
	defer teardown()
	t.Run("test redis factory method", func(t *testing.T) {
		resp := NewRedisCache(redisClient)
		assert.NotNil(t, resp, "test ok redis factory")
	})
}
func TestSetRedisValue(t *testing.T) {
	setup()
	defer teardown()
	redisKey := aws.String("redis_key")
	redisValue := aws.String("redis_value")
	ttl := 1 * time.Minute
	t.Run("test redis set value", func(t *testing.T) {
		redisSvc.SetRedisValue(redisKey, redisValue, ttl)
	})
}

func TestGetRedisValue(t *testing.T) {
	setup()
	defer teardown()
	redisKeyFailed := aws.String("redis_key_falied")
	redisKey := aws.String("redis_key")
	redisValue := aws.String("redis_value")
	ttl := 1 * time.Minute
	t.Run("test ok redis set value", func(t *testing.T) {
		redisSvc.SetRedisValue(redisKey, redisValue, ttl)
	})

	t.Run("test ok redis get value", func(t *testing.T) {
		res := redisSvc.GetRedisValue(redisKey)
		assert.Equal(t, *res, *redisValue)
	})

	t.Run("test nok redis get value", func(t *testing.T) {
		res := redisSvc.GetRedisValue(redisKeyFailed)
		assert.Nil(t, res, "test nok redis get err not nil")
	})
}
