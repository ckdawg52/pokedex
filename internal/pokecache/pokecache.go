package pokecache

import (
	"time"
	"sync"
)


type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mux *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	m := make(map[string]cacheEntry)
	newCache := Cache{cache: m, mux: &sync.Mutex{}}

	go newCache.reapLoop(interval)

	return &newCache
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	newCacheEntry := cacheEntry{createdAt: time.Now(), val: value}
	c.cache[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}

