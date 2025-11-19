package cache

import (
	"container/list"
	"sync"
)

type node[K comparable, V any] struct {
	key   K
	value V
}

type Cache[K comparable, V any] struct {
	size        uint
	elementList *list.List
	elements    map[K]*list.Element
	mutex       sync.Mutex
}

func NewCache[K comparable, V any](size uint) *Cache[K, V] {
	return &Cache[K, V]{size: size, elementList: list.New(), elements: make(map[K]*list.Element)}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.elements[key]; ok {
		elem.Value = value
		c.elementList.MoveToFront(elem)
		return
	}

	elem := c.elementList.PushFront(&node[K, V]{key: key, value: value})
	c.elements[key] = elem

	if uint(c.elementList.Len()) > c.size {
		lastElem := c.elementList.Back()
		if lastElem != nil {
			c.elementList.Remove(lastElem)
			key := lastElem.Value.(*node[K, V])
			delete(c.elements, key.key)
		}
	}
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.elements[key]; ok {
		c.elementList.MoveToFront(elem)
		return elem.Value.(*node[K, V]).value, true
	}

	var notFound V
	return notFound, false
}

func (c *Cache[K, V]) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.elements = make(map[K]*list.Element)
	c.size = 0
	c.elementList = list.New()
}
