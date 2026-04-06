package traversal

import (
	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
	"github.com/mhrdini/godsa/problems/bbgci/3_linkedlists/restructuring"
	"golang.org/x/exp/constraints"
)

/**
 * Problem:
 * Given the head of a singly linked list, determine if it is a palindrome.
 */

/**
 * Solution:
 *	- Get midpoint
 *	- Compare first half with reverse of second half
 *
 * Time: O(n)
 * Space: O(1)
 */

func PalindromicLinkedList[T constraints.Ordered](l *singlylinkedlist.List[T]) bool {
	var midpoint, double = l.Head(), l.Head()

	for double != nil && double.Next() != nil {
		midpoint = midpoint.Next()
		double = double.Next().Next()
	}

	reverse_head := restructuring.LinkedListReversalRecursiveHelper(midpoint)

	l1, l2 := l.Head(), reverse_head

	for l1 != nil && l2 != nil {
		v1 := l1.Value()
		v2 := l2.Value()
		if v1 != v2 {
			return false
		}
		l1 = l1.Next()
		l2 = l2.Next()
	}

	return true
}
