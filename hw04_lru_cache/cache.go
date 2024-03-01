package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type KeyValue struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	v, ok := c.items[key]
	if ok {
		v.Value = KeyValue{key: key, value: value}
		c.queue.MoveToFront(v)
	} else {
		newListItem := c.queue.PushFront(KeyValue{key: key, value: value})
		c.items[key] = newListItem
		if c.queue.Len() > c.capacity {
			last := c.queue.Back()
			c.queue.Remove(last)
			if kv, ok := last.Value.(KeyValue); ok {
				delete(c.items, kv.key)
			}
		}
		return false
	}
	return true
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	v, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(v)
		if kv, ok := v.Value.(KeyValue); ok {
			return kv.value, ok
		}
		return v.Value, ok
	}
	return nil, ok
}

func (c *lruCache) Clear() {
	for k, v := range c.items {
		delete(c.items, k)
		c.queue.Remove(v)
	}
}
