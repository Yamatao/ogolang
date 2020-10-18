package hw04_lru_cache //nolint:golint,stylecheck

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
	Len() int
}

type lruCache struct {
	capacity int

	mtx   sync.Mutex
	queue List
	items map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	item, found := c.items[key]
	if found {
		c.queue.MoveToFront(item)
		item.Value = cacheItem{key, value}
		return true
	}

	item = c.queue.PushFront(cacheItem{key, value})

	// eviction of the least-recent-unit
	if c.capacity > 0 && c.queue.Len() > c.capacity {
		value := c.queue.PopBack()
		delete(c.items, value.(cacheItem).cacheKey)
	}
	c.items[key] = item
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	item, found := c.items[key]
	if found {
		c.queue.MoveToFront(item)
		return item.Value.(cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.queue.Clear()
	c.items = make(map[Key]*ListItem)
}

func (c *lruCache) Len() int {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	return c.queue.Len()
}

type cacheItem struct {
	cacheKey Key
	value    interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		mtx:      sync.Mutex{},
		queue:    NewList(),
		items:    make(map[Key]*ListItem),
	}
}
