package cache

import "container/list"

//https://github.com/UncleBig/goCache/blob/master/goCache.go
//https://www.jianshu.com/p/970f1a8dd9cf
type Cache struct {
	MaxEntries int
	list       *list.List
	cache      map[string]*list.Element
}

// entry 表示一个缓存键值对
type entry struct {
	key   string
	value interface{}
}

// New 函数新建一个LRU缓存对象
func New(max int) *Cache {
	return &Cache{
		MaxEntries: max,
		list:       list.New(),
		cache:      make(map[string]*list.Element),
	}
}

// Put 函数添加一个缓存项到Cache对象中
func (cache *Cache) Put(key string, value interface{}) {
	if cache.cache == nil {
		cache.cache = make(map[string]*list.Element)
		cache.list = list.New()
	}
	// 如果缓存已经存在于Cache中，那么将该缓存项移到双向链表的最前端
	if value, ok := cache.cache[key]; ok {
		cache.list.MoveToFront(value)
		value.Value.(*entry).value = value
		return
	}

	// 将新添加的缓存项放入双向链表的最前端
	element := cache.list.PushFront(&entry{key, value})
	cache.cache[key] = element

	// 如果超出缓存容量，那么移除双向链表中的最后一项
	if cache.MaxEntries != 0 && cache.list.Len() > cache.MaxEntries {
		cache.RemoveOldest()
	}
}

// Get 方法获取具有指定键的缓存项
func (cache *Cache) Get(key string) (value interface{}, ok bool) {
	if cache.cache == nil {
		return
	}
	if element, hit := cache.cache[key]; hit {
		cache.list.MoveToFront(element)
		return element.Value.(*entry).value, true
	}
	return
}

// Remove 方法移除具有指定键的缓存
func (cache *Cache) Remove(key string) {
	if cache.cache == nil {
		return
	}
	if element, hit := cache.cache[key]; hit {
		cache.removeElement(element)
	}
}

// RemoveOldest 移除双向链表中访问时间最远的那一项
func (cache *Cache) RemoveOldest() {
	if cache.cache == nil {
		return
	}

	element := cache.list.Back()
	if element != nil {
		cache.removeElement(element)
	}
}

func (cache *Cache) removeElement(element *list.Element) {
	cache.list.Remove(element)
	entry := element.Value.(*entry)
	delete(cache.cache, entry.key)
}

// Len 方法获取Cache中包含的缓存项个数
func (cache *Cache) Len() int {
	if cache.cache == nil {
		return 0
	}
	return cache.list.Len()
}

// Clear 清除整个Cache对象
func (cache *Cache) Clear() {
	cache.list = nil
	cache.cache = nil
}
