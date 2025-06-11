package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
	mux   *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cacheMeOutside := Cache{
		cache: make(map[string]CacheEntry),
		mux:   &sync.Mutex{},
	}

	go cacheMeOutside.reapLoop(interval)

	return cacheMeOutside
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = CacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	ce, exists := c.cache[key]
	return ce.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(curr time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for key, val := range c.cache {
		if !curr.Before(val.createdAt.Add(last)) {
			delete(c.cache, key)
		}
	}
}
