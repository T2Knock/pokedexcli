package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	mux     *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	value     []byte
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		value:     value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	cache, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return cache.value, true
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

	for k, v := range c.entries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.entries, k)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: make(map[string]CacheEntry),
		mux:     &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}
