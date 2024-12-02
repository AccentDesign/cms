package pages

import (
	"sync"
	"time"
)

type item[V any] struct {
	value      V
	expiration int64
}

// Cache is a thread-safe in-memory cache using sync.Map.
type Cache[K comparable, V any] struct {
	items   sync.Map
	ttl     time.Duration
	cleanup time.Duration
}

func NewCache[K comparable, V any](defaultTTL, cleanupInterval time.Duration) *Cache[K, V] {
	c := &Cache[K, V]{
		ttl:     defaultTTL,
		cleanup: cleanupInterval,
	}
	if cleanupInterval > 0 {
		go c.startCleanup()
	}
	return c
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	value, ok := c.items.Load(key)
	if !ok {
		var zero V
		return zero, false
	}
	itm := value.(item[V])
	if itm.isExpired() {
		c.items.Delete(key)
		var zero V
		return zero, false
	}
	return itm.value, true
}

func (c *Cache[K, V]) Set(key K, value V, ttl time.Duration) {
	if ttl == 0 {
		ttl = c.ttl
	}
	var expiration int64
	if ttl > 0 {
		expiration = time.Now().Add(ttl).UnixNano()
	}
	itm := item[V]{
		value:      value,
		expiration: expiration,
	}
	c.items.Store(key, itm)
}

func (c *Cache[K, V]) Delete(key K) {
	c.items.Delete(key)
}

func (itm *item[V]) isExpired() bool {
	if itm.expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > itm.expiration
}

func (c *Cache[K, V]) startCleanup() {
	ticker := time.NewTicker(c.cleanup)
	defer ticker.Stop()

	for range ticker.C {
		c.deleteExpired()
	}
}

func (c *Cache[K, V]) deleteExpired() {
	c.items.Range(func(key, value any) bool {
		itm := value.(item[V])
		if itm.isExpired() {
			c.items.Delete(key)
		}
		return true
	})
}
