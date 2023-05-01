package doublylinkedlist

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/utils"
)

type List[T any] struct {
	size int
	head *node[T]
	tail *node[T]
}

type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

// New receives a variadic input of values whose type are T.
// Returns a *List of nodes whose values are vs.
func New[T any](vs ...T) *List[T] {
	list := &List[T]{}
	list.Add(vs...)
	return list
}

// Size returns the number of nodes inside the List.
func (l *List[T]) Size() int {
	return l.size
}

// Empty checks if the List has no nodes.
func (l *List[T]) Empty() bool {
	return l.size == 0 && l.head == nil && l.tail == nil
}

// Values returns a slice of the values carried by all the nodes within the List,
// in the same sequence.
func (l *List[T]) Values() []T {
	vs := make([]T, l.size)
	for i, n := 0, l.head; n != nil; i, n = i+1, n.next {
		vs[i] = n.value
	}
	return vs[:l.size:l.size]
}

// Reset clears the size, and head and tail nodes of the List (but the List itself is not nil).
func (l *List[T]) Reset() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

// Sort receives a utils.Comparator function used to sort the values in-place
func (l *List[T]) Sort(comp utils.Comparator[T]) {
	vs := l.Values()
	utils.Sort(vs, comp)
	l.Reset()
	l.Add(vs...)
}

// String returns a string representation of the List.
func (l *List[T]) String() string {
	return fmt.Sprintf("%v", l.Values())
}

// Add creates and inserts a new node at the tail end of the List, for every T value received as input.
func (l *List[T]) Add(vs ...T) bool {
	if len(vs) == 0 {
		return false
	}

	for _, v := range vs {
		newNode := &node[T]{
			value: v,
			prev:  nil,
			next:  nil,
		}
		if l.size == 0 {
			l.head = newNode
			l.tail = newNode
		} else {
			l.tail.next = newNode
			newNode.prev = l.tail
			l.tail = newNode
		}
		l.size++
	}

	return true
}

// InsertAt inserts new nodes with vs as values starting at position i, with zero-based indexing.
// Returns a boolean value determining whether there were any values inserted.
func (l *List[T]) InsertAt(i int, vs ...T) bool {
	if !l.withinRange(i) || len(vs) == 0 {
		return false
	}

	if i == l.size {
		return l.Add(vs...)
	}

	var start, end, prev *node[T]
	for idx, v := range vs {
		newNode := &node[T]{
			value: v,
			prev:  prev,
			next:  nil,
		}
		if idx == 0 {
			start = newNode
		}
		if idx == len(vs)-1 {
			end = newNode
		}
		if prev != nil {
			prev.next = newNode
		}
		prev = newNode
	}

	switch i {
	case 0:
		if l.head == nil {
			l.head = start
			l.tail = end
		} else {
			end.next = l.head
			l.head.prev = end
			l.head = start
		}
	default:
		var curr *node[T]
		if i <= l.size/2 {
			curr = l.head
			for pos := 0; pos < i-1; curr, pos = curr.next, pos+1 {
			}
		} else {
			curr = l.tail
			for pos := l.size - 1; pos >= i; curr, pos = curr.prev, pos-1 {
			}
		}

		start.prev = curr
		end.next = curr.next
		curr.next.prev = end
		curr.next = start
	}
	l.size += len(vs)
	return true
}

// Prepend inserts new nodes with vs as values at the start of the List.
// Returns a boolean value determining whether there were any values inserted.
func (l *List[T]) Prepend(vs ...T) bool {
	return l.InsertAt(0, vs...)
}

// Append inserts new nodes with vs as values at the end of the List, an alias for Add.
// Returns a boolean value determining whether there were any values inserted.
func (l *List[T]) Append(v ...T) bool {
	return l.Add(v...)
}

// Remove removes a node at position n, with zero-based indexing.
// n must be between 0 and l.size - 1, inclusive.
// Returns the value of the removed node (if any), and a boolean value determining whether a value
// was removed or not.
func (l *List[T]) Remove(i int) (T, bool) {
	var value T

	if l.size == 0 || i < 0 || i >= l.size {
		return value, false
	}

	switch i {
	case 0:
		value = l.head.value
		if l.size == 1 {
			l.head = nil
			l.tail = nil
		} else {
			l.head = l.head.next
		}
	default:
		var curr *node[T]
		if i <= l.size/2 {
			curr = l.head
			for pos := 0; pos < i; curr, pos = curr.next, pos+1 {
			}
		} else {
			curr = l.tail
			for pos := l.size - 1; pos > i; curr, pos = curr.prev, pos-1 {
			}
		}
		value = curr.value
		curr.prev.next = curr.next
		if curr == l.tail {
			l.tail = curr.prev
			l.tail.next = nil
		} else {
			curr.next.prev = curr.prev
		}
	}
	l.size--
	return value, true
}

// RemoveFront removes a node at the head position, using the Remove function.
// Returns the value of the removed node (if any), and a boolean value determining whether a value
// was removed or not.
func (l *List[T]) RemoveFront() (T, bool) {
	return l.Remove(0)
}

// RemoveBack removes a node at the tail position, using the Remove function.
// Returns the value of the removed node (if any), and a boolean value determining whether a value
// was removed or not.
func (l *List[T]) RemoveBack() (T, bool) {
	return l.Remove(l.size - 1)
}

func (l *List[T]) Set(i int, v T) bool {
	if !l.withinRange(i) {
		return false
	}

	switch l.size {
	case 0, i:
		return false
	default:
		var curr *node[T]
		if i <= l.size/2 {
			curr = l.head
			for pos := 0; pos < i; curr, pos = curr.next, pos+1 {
			}
		} else {
			curr = l.tail
			for pos := l.size - 1; pos > i; curr, pos = curr.prev, pos-1 {
			}
		}
		curr.value = v
		return true
	}
}

// Concat uses Add to create new nodes out of a variadic input of Lists and inserts each one into
// the List pointer receiver.
func (l *List[T]) Concat(ls ...*List[T]) {
	if len(ls) == 0 {
		return
	}

	for _, list := range ls {
		l.Add(list.Values()...)
	}
}

// withinRange returns true if the index argument is within the bounds of the list.
func (l *List[T]) withinRange(i int) bool {
	return i >= 0 && i <= l.size
}
