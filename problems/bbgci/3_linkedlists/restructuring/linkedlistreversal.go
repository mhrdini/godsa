package restructuring

import (
	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
)

/**
 * Problem:
 * Reverse a singly linked list.
 */

/**
 * Solution ->
 */

func LinkedListReversalIterative[T any](l *singlylinkedlist.List[T]) {
	prev_head := l.Head()
	var prev, curr *singlylinkedlist.Node[T] = nil, l.Head()
	for curr != nil {
		next := curr.Next()
		curr.SetNext(prev)
		prev = curr
		curr = next
	}
	l.SetHead(prev)
	l.SetTail(prev_head)
}

func LinkedListReversalRecursive[T any](l *singlylinkedlist.List[T]) {
	prev_head := l.Head()
	new_head := linkedListReversalRecursiveHelper(prev_head)
	l.SetHead(new_head)
	l.SetTail(prev_head)
}

func linkedListReversalRecursiveHelper[T any](head *singlylinkedlist.Node[T]) *singlylinkedlist.Node[T] {
	if head == nil || head.Next() == nil {
		return head
	}
	new_head := linkedListReversalRecursiveHelper(head.Next())
	head.Next().SetNext(head)
	head.SetNext(nil)
	return new_head
}
