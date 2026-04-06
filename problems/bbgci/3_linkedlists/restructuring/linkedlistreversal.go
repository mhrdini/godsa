package restructuring

import (
	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
)

/**
 * Problem:
 * Reverse a singly linked list.
 */

/**
 * Solution 1:
 *	- Iterative
 *	- Keep setting the current node's next pointer the previous node before it
 *		until you get to the end
 *
 * Time: O(n)
 * Space: O(1)
 */

/**
* Solution 2:
*	- Recursive
*	- Reverse the sublist starting at head.Next()
*	- In the base case head is a single node or nil
*		(for the case that the original list was a single node whose next is nil)
* - Therefore in the recursive case, the last node of the sublist will become
*		the new head, and the node at head.Next() is now the sublist's tail
* - In the calling context:
*		- The sublist tail, i.e. head.Next() must now be pointed to head,
*			therefore we do head.Next().SetNext(head)
*   - The original head, must now be the original list's tail, whose next
*			pointer must be nil, therefore we do head.SetNext(nil)
*
* Time: O(n)
* Space: O(n) from call stack
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
	new_head := LinkedListReversalRecursiveHelper(prev_head)
	l.SetHead(new_head)
	l.SetTail(prev_head)
}

func LinkedListReversalRecursiveHelper[T any](head *singlylinkedlist.Node[T]) *singlylinkedlist.Node[T] {
	if head == nil || head.Next() == nil {
		return head
	}
	new_head := LinkedListReversalRecursiveHelper(head.Next())
	head.Next().SetNext(head)
	head.SetNext(nil)
	return new_head
}
