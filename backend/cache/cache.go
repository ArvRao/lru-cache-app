package cache

import (
	"container/list"
	"sync"
	"time"
)

type entry struct {
	key        string
	value      interface{}
	expiration time.Time
}

type LRUCache struct {
	capacity     int
	cache        map[string]*list.Element
	evictionList *list.List
	lock         sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity:     capacity,
		cache:        make(map[string]*list.Element),
		evictionList: list.New(),
	}
}

func (c *LRUCache) Set(key string, value interface{}, ttl time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if element, ok := c.cache[key]; ok {
		c.evictionList.MoveToFront(element)
		element.Value.(*entry).value = value
		element.Value.(*entry).expiration = time.Now().Add(ttl)
		return
	}

	newItem := &entry{
		key:        key,
		value:      value,
		expiration: time.Now().Add(ttl),
	}
	element := c.evictionList.PushFront(newItem)
	c.cache[key] = element

	if c.evictionList.Len() > c.capacity {
		c.evict()
	}
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if element, ok := c.cache[key]; ok {
		if time.Now().After(element.Value.(*entry).expiration) {
			c.evictionList.Remove(element)
			delete(c.cache, key)
			return nil, false
		}

		c.evictionList.MoveToFront(element)
		return element.Value.(*entry).value, true
	}

	return nil, false
}

func (c *LRUCache) evict() {
	oldest := c.evictionList.Back()
	if oldest != nil {
		c.evictionList.Remove(oldest)
		delete(c.cache, oldest.Value.(*entry).key)
	}
}
