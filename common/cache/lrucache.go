package cache

import (
	lru "github.com/hashicorp/golang-lru"
)

type Cache struct {
	lc *lru.Cache
}

func New(size int) Cache {
	l, _ := lru.New(size)
	return Cache{
		lc: l,
	}
}

func (c Cache) Set(key string, value interface{}) {
	c.lc.Add(key, value)
}

func (c Cache) Contains(key string) bool {
	return c.lc.Contains(key)
}

func (c Cache) Get(key string) (interface{}, bool) {
	return c.lc.Get(key)
}
