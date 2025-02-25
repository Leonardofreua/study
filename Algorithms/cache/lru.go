// lru.go
// description : Least Recently Used (LRU) cache
// details : A Least Recently Used (LRU) cache is a type of cache algorithm used to manage memory within a computer. The LRU algorithm is designed to remove the least recently used items first when the cache reaches its limit.
// time complexity : O(1)
// space complexity : O(n)
// ref : https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU)

package cache

import (
	"container/list"
)

type Node struct {
	Data   int
	KeyPtr *list.Element
}

type LRUCache struct {
	Queue    *list.List
	Items    map[int]*Node
	Capacity int
}

func Init(capacity int) LRUCache {
	return LRUCache{
		Queue:    list.New(),
		Items:    make(map[int]*Node),
		Capacity: capacity,
	}
}

func (l *LRUCache) Put(key int, value int) {
	if item, ok := l.Items[key]; !ok {
		if l.Capacity == len(l.Items) {
			// if it's full, remove latest element from list and map
			back := l.Queue.Back()
			l.Queue.Remove(back)
			delete(l.Items, back.Value.(int))
		}
		l.Items[key] = &Node{Data: value, KeyPtr: l.Queue.PushFront(key)}
	} else {
		item.Data = value
		l.Items[key] = item
		l.Queue.MoveToFront(item.KeyPtr)
	}
}

func (l *LRUCache) Get(key int) int {
	if item, ok := l.Items[key]; ok {
		l.Queue.MoveToFront(item.KeyPtr)
		return item.Data
	}
	return -1
}
