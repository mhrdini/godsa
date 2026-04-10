package fractionalpointidentification

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
)

/**
 * Problem:
 * Given a singly linked list, find and return its middle node.
 * If there are two middle nodes, return the second one.
 */

/**
 * Solution:
 *	- Move slow pointer 1 at a time
 *	- Move fast pointer 2 at a time
 *	- Once fast pointer has no next node, it is at the tail
 *	- Thus slow pointer must be at the middle node
 *
 * Time: O(n) to traverse entire list
 * Space: O(1)
 */

func LinkedListMidpoint[T any](l *singlylinkedlist.List[T]) *singlylinkedlist.Node[T] {
	var midpoint, double = l.Head(), l.Head()

	for double != nil && double.Next() != nil {
		midpoint = midpoint.Next()
		double = double.Next()
		if double != nil {
			double = double.Next()
		}
	}

	return midpoint
}

func LinkedListMidpointExample() {
	l := singlylinkedlist.New(1, 2, 4, 7, 3)
	fmt.Println(LinkedListMidpoint(l))
}
