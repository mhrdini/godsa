package cycledetection

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
)

/**
 * Problem:
 * Given a singly linked list, determine if it contains a cycle.
 * A cycle occurs if a node's next pointer references an earlier node
 * in the list, causing a loop.
 */

/**
 * Solution:
 *	- Floyd's Detection Algorithm
 *	- Move slow pointer 1 at a time
 *	- Move fast pointer 2 at a time
 *	- Keep moving each until:
 *		- Fast pointer becomes null -> no cycle; or,
 *		- Fast pointer == slow pointer -> has cycle
 *  - The maximal number of steps required for the fast pointer to catch up to
 *		the slow pointer is k steps (once both are in the cycle), where k is the
 *		length of the cycle
 *		- Worst case: Cycle length = Linked list size ->  k = n
 *
 * Time: O(n)
 * Space: O(1)
 */

func LinkedListLoop[T any](l *singlylinkedlist.List[T]) bool {
	not_first := false
	for slow, fast := l.Head(), l.Head(); fast != nil; not_first = true {
		if not_first && fast == slow {
			return true
		}
		slow = slow.Next()
		fast = fast.Next()
		if fast != nil {
			fast = fast.Next()
		}
	}
	return false
}

func LinkedListLoopExample() {
	l := singlylinkedlist.New(0, 1, 2)

	// comment out to remove cycle
	cycle := l.Tail()

	l.Add(3, 4, 5)

	// comment out to remove cycle
	l.Tail().SetNext(cycle)

	fmt.Println(LinkedListLoop(l))
}
