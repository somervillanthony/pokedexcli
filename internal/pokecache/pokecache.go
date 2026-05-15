package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache:    map[string]cacheEntry{},
		mu:       sync.Mutex{},
		interval: interval,
	}
	ticker := time.NewTicker(c.interval)

	go func() {
		defer ticker.Stop()
		for range ticker.C {
			c.reapLoop()
		}
	}()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	fmt.Println("below is val []bytes")
	fmt.Println(entry.val)
	c.cache[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cache[key]
	if ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, val := range c.cache {
		elapsed := time.Since(val.createdAt)
		if elapsed >= c.interval {
			delete(c.cache, key)
		}
	}
}
