package heapsort

import (
	"github.com/mhrdini/godsa/datastructures/lists/arraylist"
	"github.com/mhrdini/godsa/datastructures/trees/heaps/binaryheap"
	"github.com/mhrdini/godsa/datastructures/utils/comparator"
	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](s []T, ascending bool) []T {
	var heap *binaryheap.Heap[T]
	if ascending {
		heap = binaryheap.MinHeap(comparator.OrderedComparator[T], s...)
	} else {
		heap = binaryheap.MaxHeap(comparator.OrderedComparator[T], s...)
	}

	result := arraylist.New[T]()
	for !heap.Empty() {
		v, ok := heap.Pop()
		if ok {
			result.Add(v)
		}
	}
	return result.Values()
}
