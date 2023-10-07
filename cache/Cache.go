package cache

import (
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c Cache) Add(key string, val []byte) {
	c.cacheMap[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	cache, ok := c.cacheMap[key]
	if !ok {
		return nil, false
	} else {
		return cache.val, true
	}
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheMap: map[string]cacheEntry{},
	}
	go func() {
		for range time.NewTicker(1 * time.Second).C {
			for k, ce := range c.cacheMap {
				if time.Now().UnixMilli()-ce.createdAt.UnixMilli() >= interval.Milliseconds() {
					delete(c.cacheMap, k)
				}
			}
		}
	}()
	return c
}
