// lru.go
// description : Least Recently Used (LRU) cache
// details : A Least Recently Used (LRU) cache is a type of cache algorithm used to manage memory within a computer. The LRU algorithm is designed to remove the least recently used items first when the cache reaches its limit.
// time complexity : O(1)
// space complexity : O(n)
// ref : https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU)

package cache

import "container/list"

type LRUNode[T any] struct {
	Data   T
	KeyPtr *list.Element
}

type LRU[T any] struct {
	Queue    *list.List
	Items    map[int]*LRUNode[T]
	Capacity int
}

func InitLRU[T any](capacity int) LRU[T] {
	return LRU[T]{
		Queue:    list.New(),
		Items:    make(map[int]*LRUNode[T]),
		Capacity: capacity,
	}
}

func (l *LRU[T]) Put(key int, value T) {
	if item, ok := l.Items[key]; !ok {
		if l.Capacity == len(l.Items) {
			back := l.Queue.Back()
			l.Queue.Remove(back)
			delete(l.Items, back.Value.(int))
		}
		l.Items[key] = &LRUNode[T]{Data: value, KeyPtr: l.Queue.PushFront(key)}
	} else {
		item.Data = value
		l.Items[key] = item
		l.Queue.MoveToFront(item.KeyPtr)
	}
}

func (l *LRU[T]) Get(key int) T {
	if item, ok := l.Items[key]; ok {
		l.Queue.MoveToFront(item.KeyPtr)
		return item.Data
	}

	var zero T
	return zero
}
