package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Returns new cache and starts a reapLoop
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

// When called, deletes entries that are older than interval
// Should be used in a loop with a ticker
func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()

		for key, entry := range c.cache {
			if time.Now().After(entry.createdAt.Add(interval)) {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
