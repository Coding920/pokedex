package pokecache

import (
	"time"
)

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	return
}

// Returns nil and false if key not in cache
// or returns the byte slice in cacheEntry.val and true
func (c Cache) Get(key string) (value []byte, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}
