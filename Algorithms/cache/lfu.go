// description: LFU (Least Frequently Used) Cache is a caching algorithm where the least frequently accessed cache block is removed when the cache reaches its capacity. In LFU, we take into account how often a page is accessed and how recently it was accessed. If a page has a higher frequency than another, it cannot be removed, as it is accessed more frequently if multiple pages share the same frequency, the one that was accessed least recently (the Least Recently Used, or LRU page) is removed. This is typically handled using a FIFO (First-In-First-Out) method, where the oldest page among those with the same frequency is the one to be removed.
// ref: (https://en.wikipedia.org/wiki/Least_frequently_used)
// time complexity: O(N)
// space complexity: O(1)

package cache

import (
	"container/list"
)

const FIRST_FREQUENCY int = 1

type LFUNode struct {
	Value     any
	Key       int
	Frequency int
}

type LFU struct {
	// key: key of element
	// value: element
	Items map[int]*list.Element

	// key: frequency of possible occurrences of all elements in the Items
	// value: elements with the same frequency
	Frequency map[int]*list.List

	Capacity int
	MinFreq  int // Tracks the minimum frequency
}

func InitLFU(capacity int) LFU {
	return LFU{
		Items:     make(map[int]*list.Element),
		Frequency: make(map[int]*list.List),
		Capacity:  capacity,
		MinFreq:   0,
	}
}

func (l *LFU) Get(key int) any {
	if node, ok := l.Items[key]; ok {
		l.updateNodeFrequency(node)
		return node.Value.(*LFUNode).Value
	}

	var zero any
	return zero
}

func (l *LFU) Put(key int, value any) {
	// Do nothing if the cache has zero capacity
	if l.Capacity == 0 {
		return
	}

	if node, ok := l.Items[key]; ok {
		node.Value.(*LFUNode).Value = value
		l.updateNodeFrequency(node)
	} else {
		// Remove the least frequently used node if cache is full
		if l.Capacity == len(l.Items) {
			minFreq := l.Frequency[l.MinFreq]
			backMinFreq := minFreq.Back()

			minFreq.Remove(backMinFreq)
			delete(l.Items, backMinFreq.Value.(*LFUNode).Key)
		}

		newNode := &LFUNode{Value: value, Key: key, Frequency: FIRST_FREQUENCY}
		l.insert(newNode)
		l.MinFreq = FIRST_FREQUENCY
	}
}

// Add a node right after the head
func (l *LFU) insert(node *LFUNode) {
	nodeFreq, ok := l.Frequency[node.Frequency]

	if !ok { // Initialize the frequency list if it doesn't exist
		nodeFreq = list.New()
		l.Frequency[node.Frequency] = nodeFreq
	}
	l.Items[node.Key] = nodeFreq.PushFront(node)
}

func (l *LFU) updateNodeFrequency(node *list.Element) {
	nodeObj := node.Value.(*LFUNode)
	oldFreq := l.Frequency[nodeObj.Frequency]

	nodeObj.Frequency++
	oldFreq.Remove(node)

	if l.MinFreq == nodeObj.Frequency && oldFreq.Len() == 0 {
		// If it is the last node of the minimum frequency that is removed
		l.MinFreq++
	}
	l.insert(nodeObj)
}
