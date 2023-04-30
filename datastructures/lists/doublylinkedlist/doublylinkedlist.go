package doublylinkedlist

type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type List[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (l *List[T]) Iterate() (vs []T) {
	for n := l.head; n != nil; n = n.Next() {
		vs = append(vs, n.value)
	}
	return
}

// InsertAt inserts a new node with value v in position n, using zero-based indexing
func (l *List[T]) InsertAt(v T, n int) bool {
	if n > l.size || n < 0 {
		return false
	}

	node := &Node[T]{
		value: v,
		prev:  nil,
		next:  nil,
	}

	switch n {
	case 0:
		if l.size == 0 {
			l.head = node
			l.tail = node
		} else {
			l.head.prev = node
			node.next = l.head
			l.head = node
		}
	case l.size:
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	default:
		var curr *Node[T]
		if n <= l.size/2 {
			curr = l.head
			for pos := 0; pos <= n-1; curr, pos = curr.next, pos+1 {
			}
		} else {
			curr = l.tail
			for pos := l.size - 1; pos >= n+1; curr, pos = curr.prev, pos-1 {
			}
		}
		if curr == l.tail {
			l.tail = node
		}
		node.next = curr.next
		node.prev = curr
		curr.next = node
	}
	l.size++
	return true
}

func (l *List[T]) InsertFront(v T) bool {
	return l.InsertAt(v, 0)
}

func (l *List[T]) InsertBack(v T) bool {
	return l.InsertAt(v, l.size)
}

// Remove removes a node at position n, with zero-based indexing.
// n must be between 0 and l.size - 1, inclusive.
func (l *List[T]) Remove(n int) (T, bool) {
	var value T

	if l.size == 0 || n >= l.size || n < 0 {
		return value, false
	}

	switch n {
	case 0:
		value = l.head.value
		if l.size == 1 {
			l.head = nil
			l.tail = nil
		} else {
			l.head = l.head.next
		}
	default:
		var curr *Node[T]
		if n <= l.size/2 {
			curr = l.head
			for pos := 0; pos < n; curr, pos = curr.next, pos+1 {
			}
		} else {
			curr = l.tail
			for pos := l.size - 1; pos > n; curr, pos = curr.prev, pos-1 {
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

func (l *List[T]) RemoveFront() (T, bool) {
	return l.Remove(0)
}

func (l *List[T]) RemoveBack() (T, bool) {
	return l.Remove(l.size - 1)
}

func (l *List[T]) ConcatUnsafe(ls ...*List[T]) {
	if len(ls) == 0 {
		return
	}

	for _, list := range ls {
		if l.size == 0 {
			l.size = list.size
			l.head = list.head
			l.tail = list.tail
		} else if list.size != 0 {
			list.head.prev = l.tail
			l.tail.next = list.head
			l.tail = list.tail
			l.size += list.size
		}
	}
}
