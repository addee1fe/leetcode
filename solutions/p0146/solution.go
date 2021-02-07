package solution

import "container/list"

// main idea of how to create LRU Cache took from HashiCorp repo
// https://github.com/hashicorp/golang-lru
// so technically this solution is not mine

// LRUCache ...
type LRUCache struct {
	size      int
	evictList *list.List
	items     map[int]*list.Element
}

// used to hold a value in the evictList
type entry struct {
	key   int
	value int
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	return LRUCache{
		size:      capacity,
		evictList: list.New(),
		items:     make(map[int]*list.Element),
	}
}

// Get returns -1 if there is no enry in cache
func (l *LRUCache) Get(key int) int {
	if l.items == nil {
		return -1
	}
	if ele, hit := l.items[key]; hit {
		l.evictList.MoveToFront(ele)
		return ele.Value.(*entry).value
	}
	return -1
}

// Put ...
func (l *LRUCache) Put(key int, value int) {
	if l.items == nil {
		return
	}
	if ee, ok := l.items[key]; ok {
		l.evictList.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}
	ele := l.evictList.PushFront(&entry{key, value})
	l.items[key] = ele
	if l.evictList.Len() > l.size {
		l.removeOldest()
	}
}

func (l *LRUCache) removeOldest() {
	ent := l.evictList.Back()
	if ent != nil {
		l.removeElement(ent)
	}
}

// removeElement is used to remove a given list element from the cache
func (l *LRUCache) removeElement(e *list.Element) {
	l.evictList.Remove(e)
	kv := e.Value.(*entry)
	delete(l.items, kv.key)
}
