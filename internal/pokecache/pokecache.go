package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		cache: map[string]cacheEntry{},
		mu:    sync.Mutex{},
	}
	ticker := time.NewTicker(interval * time.Second)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			newCache.reapLoop()
		}
	}()
	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
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

func (c *Cache) reapLoop()
