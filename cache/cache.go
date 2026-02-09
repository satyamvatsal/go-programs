package cache

import (
	"container/list"

	"github.com/satyamvatsal/go-programs/pair"
)

type Cache[T any] interface {
	Get(key string) T
	Put(key string, value T)
	evict()
}
type LRUCache[T any] struct {
	dll      *list.List
	capacity int
	cacheMap map[string]*list.Element
}

func NewCache[T any](capacity int) Cache[T] {
	return &LRUCache[T]{
		dll:      list.New(),
		cacheMap: make(map[string]*list.Element, capacity),
		capacity: capacity,
	}
}
func (c *LRUCache[T]) Get(key string) T {
	node, ok := c.cacheMap[key]
	if !ok {
		var zero T
		return zero
	}
	p := node.Value.(pair.Pair[string, T])
	c.dll.MoveToFront(node)
	return p.Second
}
func (c *LRUCache[T]) evict() {
	// lru eviction
	lru := c.dll.Back()
	p := lru.Value.(pair.Pair[string, T])
	key := p.First
	c.dll.Remove(c.cacheMap[key])
	delete(c.cacheMap, key)
}
func (c *LRUCache[T]) Put(key string, value T) {
	node, ok := c.cacheMap[key]
	if ok {
		delete(c.cacheMap, key)
		c.dll.Remove(node)
	} else if len(c.cacheMap) >= c.capacity {
		c.evict()
	}
	p := pair.NewPair(key, value)
	var mru *list.Element
	if c.dll.Front() == nil {
		mru = c.dll.PushBack(p)
	} else {
		mru = c.dll.InsertBefore(p, c.dll.Front())
	}
	c.cacheMap[key] = mru
}
