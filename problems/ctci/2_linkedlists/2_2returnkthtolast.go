package linkedlists

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
)

// Implement an algorithm to find the kth to last element of a singly linked list.

func ReturnKthToLast[T any](list *singlylinkedlist.List[T], k int) (*singlylinkedlist.Node[T], error) {
	iter := list.NewIterator()
	var current, kBehind *singlylinkedlist.Node[T]

	for i := 0; i < k; i++ {
		if n, ok := iter.Next(); ok {
			current = n
		} else {
			return kBehind, fmt.Errorf("error: k-th to last element does not exist")
		}
	}

	iter.Reset()
	kBehind, _ = iter.Next()
	for !current.Next().Empty() {
		current = current.Next()
		kBehind, _ = iter.Next()
	}
	return kBehind, nil
}
