package pokecache

import (
	"fmt"
	"time"
)

// WARNING: something is funky with the mutex and the warning, works fine but I still don't like it
func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entry[key] = newEntry
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if value, ok := c.entry[key]; !ok {
		return []byte{}, false
	} else {
		return value.val, true
	}
}

func (c Cache) reapLoop(interval time.Duration) {
	// NOTE: this works but should probably still use a ticker, should doublecheck
	for {
		time.Sleep(interval)

		c.mu.Lock()
		for key, value := range c.entry {
			if time.Since(value.createdAt) > interval {
				// fmt.Println("An entry has been reaped from the cache")
				delete(c.entry, key)
			}
		}

		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{}
	newCache.entry = make(map[string]cacheEntry)
	go newCache.reapLoop(interval)

	return &newCache
}
