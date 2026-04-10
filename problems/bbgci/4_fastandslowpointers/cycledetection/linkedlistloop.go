package cycledetection

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
)

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
