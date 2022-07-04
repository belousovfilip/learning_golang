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

func (l *lruCache) Set(key Key, value interface{}) bool {
	element, exists := l.items[key]
	if exists == true {
		l.queue.MoveToFront(element)
		element.Value.(*CacheItem).value = value
		return true
	}
	if l.queue.Len() == l.capacity {
		listItem := l.queue.Back()
		l.queue.Remove(listItem)
		delete(l.items, listItem.Value.(*CacheItem).key)
	}
	item := NewCacheItem(key, value)
	element = l.queue.PushFront(item)
	l.items[item.key] = element
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	element, exists := l.items[key]
	if exists == false {
		return nil, false
	}
	l.queue.MoveToFront(element)
	return element.Value.(*CacheItem).value, true
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

type CacheItem struct {
	key   Key
	value interface{}
}

func NewCacheItem(key Key, value interface{}) *CacheItem {
	return &CacheItem{
		key:   key,
		value: value,
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
