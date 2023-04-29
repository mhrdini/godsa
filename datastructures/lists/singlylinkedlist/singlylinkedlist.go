package singlylinkedlist

type Node[T any] struct {
	value T
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

// InsertAt inserts a new node with value v at position n, with zero-based indexing.
func (l *List[T]) InsertAt(v T, n int) bool {
	if n > l.size || n < 0 {
		return false
	}

	node := &Node[T]{
		value: v,
		next:  nil,
	}

	switch n {
	case 0:
		if l.size == 0 {
			l.head = node
			l.tail = node
		} else {
			node.next = l.head
			l.head = node
		}
	case l.size:
		l.tail.next = node
		l.tail = node
	default:
		curr := l.head
		for pos := 0; pos <= n-1; curr, pos = curr.next, pos+1 {
		}
		if curr == l.tail {
			l.tail = node
		}
		node.next = curr.next
		curr.next = node
	}
	l.size++
	return true
}

func (l *List[T]) Push(v T) bool {
	return l.InsertAt(v, 0)
}

func (l *List[T]) Add(v T) bool {
	return l.InsertAt(v, l.size)
}

func (l *List[T]) Append(v T) bool {
	return l.InsertAt(v, l.size)
}

// DeleteAt removes a node at position n, with zero-based indexing.
// n must be between 0 and l.size - 1, inclusive.
func (l *List[T]) DeleteAt(n int) (T, bool) {

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
		var prev *Node[T]
		for curr, pos := l.head, 0; pos <= n; curr, pos = curr.next, pos+1 {
			if pos == n {
				value = curr.value
				prev.next = curr.next
				if curr == l.tail {
					l.tail = prev
				}
			}
			prev = curr
		}
	}

	l.size--
	return value, true
}

func (l *List[T]) DeleteFront() (T, bool) {
	return l.DeleteAt(0)
}

func (l *List[T]) DeleteBack() (T, bool) {
	return l.DeleteAt(l.size - 1)
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
			l.tail.next = list.head
			l.tail = list.tail
			l.size += list.size
		}
	}
}
