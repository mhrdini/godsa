package main

import (
	"github.com/mhrdini/godsa/datastructures/trees/heaps/binomialheap"
	"github.com/mhrdini/godsa/datastructures/utils/comparator"
)

func main() {
	h := binomialheap.MinHeap(comparator.OrderedComparator[int], 3)
	h.Insert(183)
	h.Insert(95)
	h.Insert(117)
	h.Insert(44)

	// h.Insert(5)
	// h.Insert(3)
	// h.Insert(6)
	// h.Insert(4)
	// h.Insert(8)

	// h.Insert(2)
	// h.Insert(1)
	// h.Insert(6)
	// fmt.Println(h.Root())
}
