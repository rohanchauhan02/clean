package cache

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

// QoalaCache will return object of implementation in NewGoCache
type QoalaCache struct {
	c *cache.Cache
}

// NewGoCache will return implementation of qoala go-cache
func NewGoCache() QoalaCache {
	return QoalaCache{
		c: cache.New(5*time.Minute, 10*time.Minute),
	}
}

// Set method go-cache to set data into memory
func (qc QoalaCache) Set(key string, value interface{}, duration time.Duration) {
	qc.c.Set(key, value, duration)
}

// Get method go-cache to get data from memory with given key
func (qc QoalaCache) Get(key string) (interface{}, bool) {
	return qc.c.Get(key)
}

// Delete method go-cache to delete data from memory with given key
func (qc QoalaCache) Delete(key string) {
	qc.c.Delete(key)
}
