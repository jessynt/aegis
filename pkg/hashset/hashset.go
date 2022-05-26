package hashset

import (
	"sync"
)

type Set struct {
	mu    sync.RWMutex
	items map[interface{}]struct{}
}

func New(values ...interface{}) *Set {
	set := &Set{
		items: make(map[interface{}]struct{}),
	}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

func (c *Set) Add(items ...interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, item := range items {
		c.items[item] = struct{}{}
	}
}

func (c *Set) contains(item interface{}) bool {
	_, found := c.items[item]
	return found
}

func (c *Set) Contains(item interface{}) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.contains(item)
}

func (c *Set) Remove(items ...interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, item := range items {
		delete(c.items, item)
	}
}

func (c *Set) Size() int {
	return len(c.items)
}

func (c *Set) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[interface{}]struct{})
}

func (c *Set) Values() []interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	values := make([]interface{}, c.Size())
	count := 0

	for item := range c.items {
		values[count] = item
		count++
	}
	return values
}
