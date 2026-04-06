package restructuring

import "github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"

/**
 * Problem:
 * Return the head of a singly linked list after removing the k-th node
 * from the end of it.
 *
 * Constraints:
 *	- The linked list contains at least one node.
 */

func RemoveKthLastNode[T any](l *singlylinkedlist.List[T], k int) {
	var dummy = &singlylinkedlist.Node[T]{}
	dummy.SetNext(l.Head())

	// left -> node before k-th last
	// right -> the node that is k nodes ahead of left, meant to finally be the
	// last node in the list
	var left, right *singlylinkedlist.Node[T] = dummy, dummy
	for range k {
		right = right.Next()
	}
	for right != nil && right.Next() != nil {
		left = left.Next()
		right = right.Next()
	}
	removed := left.Next()
	left.SetNext(removed.Next())
	if removed == l.Head() {
		l.SetHead(left.Next())
	}
	if removed == l.Tail() {
		l.SetTail(left)
	}
	l.SetSize(l.Size() - 1)
}
