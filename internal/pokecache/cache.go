package pokecache

import(
	"time"
	"sync"
)
type cacheEntry struct{
	createdAt 	time.Time
	val			[]byte
}

type Cache struct{
	entry map[string]cacheEntry
	mu sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entry: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entry[key] = cacheEntry{
		createdAt: 	time.Now(),
		val: 		val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	_, ok := c.entry[key]
	if ok {
		return c.entry[key].val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for i := range c.entry {
			if time.Since(c.entry[i].createdAt) > interval {
				delete(c.entry, i)
			}
		}
		c.mu.Unlock()
	}
}
