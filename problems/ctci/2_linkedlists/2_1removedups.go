package linkedlists

import (
	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
)

// Write code to remove duplicates from an unsorted linked list.

func RemoveDups[T comparable](list *singlylinkedlist.List[T]) {
	seen := make(map[T]bool)
	iter := list.NewIterator()
	for v, ok := iter.Next(); ok; {
		dup, exists := seen[v.Value()]
		if exists && dup {
			v, ok = iter.Next()
			list.Remove(iter.Index() - 1)
			continue
		} else if !dup {
			seen[v.Value()] = true
		}
		v, ok = iter.Next()
	}
}
