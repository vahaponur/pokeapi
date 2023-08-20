package pokeapi

import (
	"fmt"
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

func NewCache(duration int32) *Cache {
	cache := Cache{entiries: make(map[string]cacheEntry)}

	go readLoop(&cache, duration)

	return &cache
}
func readLoop(cache *Cache, duration int32) {
	ticker := time.NewTicker(time.Second * time.Duration(duration))
	for range ticker.C {
		newEntries := map[string]cacheEntry{}

		for key, entry := range cache.entiries {
			dif := time.Now().Sub(entry.createdAt)
			if dif > time.Second*time.Duration(duration) {
				fmt.Println("Entry passed")
				continue
			}
			fmt.Print("Entry not passed")
			newEntries[key] = entry

		}
		cache.mu.Lock()
		cache.entiries = newEntries
		cache.mu.Unlock()
	}

}
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheEntry := cacheEntry{createdAt: time.Now(), val: val}
	c.entiries[key] = cacheEntry
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	ce, ok := c.entiries[key]
	if !ok {
		return nil, false
	}
	return ce.val, true
}
