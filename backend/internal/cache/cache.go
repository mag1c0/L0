package cache

import (
	"sync"
	"time"
)

const (
	DefaultExpire time.Duration = 0
)

type item struct {
	val    interface{}
	expire int64
}

type cache struct {
	mu    sync.RWMutex
	items map[string]*item
}

type Cache struct {
	*cache
}

func New() *Cache {
	return &Cache{
		cache: &cache{
			items: make(map[string]*item),
		},
	}
}

func (c *Cache) Set(key string, val interface{}, ttl time.Duration) {
	var t int64

	if ttl == DefaultExpire {
		t = int64(DefaultExpire)
	} else if ttl > 0 {
		t = time.Now().Add(ttl).UnixNano()
	}

	c.mu.Lock()
	c.items[key] = &item{
		val:    val,
		expire: t,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	cacheItem, exist := c.items[key]
	c.mu.RUnlock()

	if exist {
		if cacheItem.expire > 0 && time.Now().UnixNano() > cacheItem.expire {
			return nil, false
		}

		return cacheItem.val, true
	}

	return nil, false
}
