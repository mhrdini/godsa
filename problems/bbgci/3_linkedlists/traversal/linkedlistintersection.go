package traversal

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
)

/**
 * Problem:
 * Return the node where two singly linked lists intersect.
 * If the linked lists don't intersect, return null.
 */

/**
 * Solution:
 *	- Use two node pointers, each going through one of the lists and checking
 *		if its next node is the node of the other pointer
 *
 * Time: O(m + n) where m, n are the lists' sizes
 * Space: O(1)
 */

func LinkedListIntersection[T any](l1, l2 *singlylinkedlist.List[T]) *singlylinkedlist.Node[T] {
	p1, p2 := l1.Head(), l2.Head()

	for p1 != nil || p2 != nil {
		if p1.Next() != nil {
			if p1.Next() != p2 {
				p1 = p1.Next()
			} else {
				return p2
			}
		}
		if p2.Next() != nil {
			if p2.Next() != p2 {
				p2 = p2.Next()
			} else {
				return p1
			}
		}
	}

	return nil
}

func LinkedListIntersectionExample() {
	suffix := singlylinkedlist.New(8, 7, 2)
	l1 := singlylinkedlist.New(1, 3, 4)
	l2 := singlylinkedlist.New(6, 4)

	l1.Tail().SetNext(suffix.Head())
	l1.SetTail(suffix.Tail())
	l1.SetSize(l1.Size() + suffix.Size())

	l2.Tail().SetNext(suffix.Head())
	l2.SetTail(suffix.Tail())
	l2.SetSize(l2.Size() + suffix.Size())

	fmt.Println("l1:", l1)
	fmt.Println("l2:", l2)
	fmt.Println("intersection:", LinkedListIntersection(l1, l2))
}
