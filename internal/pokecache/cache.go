package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	CacheEntry map[string]cacheEntry
	mu         *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		CacheEntry: map[string]cacheEntry{}, mu: &sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.CacheEntry[key] = cacheEntry{createdAt: time.Now().UTC(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.CacheEntry[key]
	return entry.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.CacheEntry {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.CacheEntry, k)
		}
	}
}
