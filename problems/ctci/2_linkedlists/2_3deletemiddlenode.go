package linkedlists

import "github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"

// Implement an algorithm to delete a node in the middle (i.e., any node but the
// first and last node, not necessarily the exact middle) of a singly linked
// list, given only access to that node.

func DeleteMiddleNode[T comparable](list *singlylinkedlist.List[T], middle T) {
	iter := list.NewIterator()
	for n, ok := iter.Next(); ok; {
		if iter.Index() != 0 && iter.Index() != list.Size() && n.Value() == middle {
			n, ok = iter.Next()
			list.Remove(iter.Index() - 1)
			continue
		}
		n, ok = iter.Next()
	}
}
