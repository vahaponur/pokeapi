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
	entiries map[string]cacheEntry
	mu       sync.RWMutex
}

func NewCache(duration time.Duration) *Cache {
	cache := Cache{entiries: make(map[string]cacheEntry)}
	defer cache.mu.RUnlock()
	cache.mu.RLock()

	go readLoop(&cache, duration)

	return &cache
}
func readLoop(cache *Cache, duration time.Duration) {
	ticker := time.NewTicker(duration * time.Second)
	for range ticker.C {
		newEntries := map[string]cacheEntry{}

		for key, entry := range cache.entiries {
			dif := time.Now().Sub(entry.createdAt)
			if dif > duration {
				continue
			}
			newEntries[key] = entry

		}
		cache.mu.Lock()
		cache.entiries = newEntries
		cache.mu.Unlock()
	}

}
func (c *Cache) Add(key string, val []byte) {
	cacheEntry := cacheEntry{createdAt: time.Now(), val: val}
	c.entiries[key] = cacheEntry
}
func (c *Cache) Get(key string) (cacheEntry, bool) {
	val, ok := c.entiries[key]
	if !ok {
		return cacheEntry{}, false
	}
	return val, true
}
