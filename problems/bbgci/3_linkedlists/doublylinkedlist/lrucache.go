package doublylinkedlist

import "fmt"

/**
 * Problem:
 * Design and implement a data structure for the Least Recently Used (LRU) cache
 * that supports the following operations:
 *	- LRUCache(capacity: int): Initialize an LRU cache with the specified
 *		capacity.
 *	- get(key: int) -> int: Return the value associated with a key. Return -1 if
 *		the key doesn't exist.
 *	- put(key: int, value: int) -> None: Add a key and its value to the cache.
 *		If adding the key would result in the cache exceeding its size capacity,
 *   evict the least recently used element. If the key already exists in the
 *		cache, update its value.
 *
 * Constraints:
 *	- All keys and values are positive integers.
 *	- The cache capacity is positive.
 *
 * Example:
 *  - Input: [put(1, 100), put(2, 250), get(2), put(4, 300), put(3, 200),
 *		get(4), get(1)], capacity = 3
 *	- Output: [250, 300, -1]
 */

/**
 * Solution:
 *  - FIFO-ish, but performing get(key int) will move the existing item with
 *		that key to the most recently used end (head)
 *  - Performing put(key, value int) while  at full capacity will remove the
 *    item at the least recently used end (tail) and add the item to the most
 *		recently used end (head)
 */

type LRUCache struct {
	capacity int
	size     int
	head     *LRUNode
	tail     *LRUNode
	nodes    map[int]*LRUNode // key -> node, ref to node
}

type LRUNode struct {
	key   int
	value int
	prev  *LRUNode
	next  *LRUNode
}

func Constructor(capacity int) *LRUCache {
	lc := &LRUCache{
		capacity: capacity,
		size:     0,
		head:     nil,
		tail:     nil,
		nodes:    make(map[int]*LRUNode, 0),
	}
	return lc
}

func (lc *LRUCache) get(key int) int {
	if n, ok := lc.nodes[key]; ok {
		prev_head := lc.head
		n_prev := n.prev
		n_next := n.next
		if n.prev != nil {
			n.prev.next = prev_head
		}
		if prev_head.next != nil {
			n.next = prev_head.next
		}
		n.prev = prev_head.prev
		n.next = prev_head.next
		prev_head.next = n_prev
		prev_head.prev = n_next
		return n.value
	}

	return -1
}

func (lc *LRUCache) put(key, value int) {
	if lc.size == lc.capacity {
		delete(lc.nodes, lc.tail.key)
		lc.tail.prev.next = nil
		lc.tail = lc.tail.prev
	}
	n := &LRUNode{
		key:   key,
		value: value,
		prev:  nil,
		next:  lc.head,
	}
	if lc.head != nil {
		lc.head.prev = n
	}
	if lc.tail == nil || lc.size == 1 {
		lc.tail = lc.head
	}
	lc.head = n
	lc.nodes[key] = n
	lc.size++
}

func LRUCacheExample() {
	var results []int

	capacity := 3
	lc := Constructor(capacity)

	lc.put(1, 100)
	lc.put(2, 250)
	results = append(results, lc.get(2))
	lc.put(4, 300)
	lc.put(3, 200)
	results = append(results, lc.get(4))
	results = append(results, lc.get(1))

	fmt.Println(results)
}
