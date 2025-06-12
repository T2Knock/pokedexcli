package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	mu      sync.Mutex
	done    chan struct{}
}

type CacheEntry struct {
	createdAt time.Time
	value     []byte
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		value:     value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cache, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return cache.value, true
}

func (c *Cache) Close() {
	close(c.done)
}

func (c *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for key, cacheEntry := range c.entries {
				if time.Since(cacheEntry.createdAt) > duration {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		case <-c.done:
			return
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]CacheEntry),
		done:    make(chan struct{}),
	}

	go c.reapLoop(interval)

	return c
}
