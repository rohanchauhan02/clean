package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	cacheKey      = "redis_key"
	cacheValue    = "redis_value"
	cacheDuration = 10 * time.Second
)

func TestNewGoCache(t *testing.T) {
	t.Run("tes NewGoCache factory method", func(t *testing.T) {
		resp := NewGoCache()
		assert.NotNil(t, resp, "test ok init go cache client")
	})
}

func TestSet(t *testing.T) {
	t.Run("test go cache set key", func(t *testing.T) {
		resp := NewGoCache()
		resp.Set(cacheKey, cacheValue, cacheDuration)
	})
}

func TestGet(t *testing.T) {
	t.Run("test go cache get key", func(t *testing.T) {
		resp := NewGoCache()
		resp.Set(cacheKey, cacheValue, cacheDuration)

		res, ok := resp.Get(cacheKey)
		assert.Equal(t, ok, true, "test ok true")
		assert.Equal(t, res, cacheValue, "test ok value match")
	})
}

func TestDelete(t *testing.T) {
	t.Run("test go cache delete key", func(t *testing.T) {
		resp := NewGoCache()
		resp.Set(cacheKey, cacheValue, cacheDuration)
		resp.Delete(cacheKey)
	})
}
