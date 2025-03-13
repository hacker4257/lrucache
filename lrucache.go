package lrucache

import (
	"container/list"
	"errors"
	"sync"
)

// LruCache represents a least-recently-used (LRU) cache.
type LruCache struct {
	queue    *list.List               // most recently used would be at the front of the queue
	items    map[string]*list.Element // map of keys to queue elements
	capacity int                      // max number of items before eviction
	mu       sync.RWMutex             // mutex to protect the cache
}

type Node struct {
	data any // The cached data.
	key  string
}

// NewLruCache creates a new LruCache with the specified capacity.
func NewLruCache(capacity int) *LruCache {
	return &LruCache{
		queue:    list.New(),
		items:    make(map[string]*list.Element),
		capacity: capacity,
	}
}

// Retrieves a cached item by key and updates its position in the cache.
func (c *LruCache) Get(key string) (any, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if item, ok := c.items[key]; ok {
		itemNode := item.Value.(Node)
		c.queue.MoveToFront(item)
		return itemNode.data, nil
	}

	return nil, errors.New("cache not found")
}

// Adds or updates a cached item with the given key and value.
func (c *LruCache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, ok := c.items[key]; !ok {
		if c.capacity == len(c.items) {
			back := c.queue.Back()
			c.queue.Remove(back)
			if selectedNode, ok := back.Value.(Node); ok {
				delete(c.items, selectedNode.key) // Retrieve the key for removal from the map.
			}
		}
		c.items[key] = c.queue.PushFront(Node{data: value, key: key})
	} else {
		item.Value = Node{data: value, key: key}
		c.items[key] = item
		c.queue.MoveToFront(item)
	}
}
