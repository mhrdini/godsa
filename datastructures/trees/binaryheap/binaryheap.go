package binaryheap

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/lists/arraylist"
	"github.com/mhrdini/godsa/datastructures/utils/comparator"
)

const (
	maxHeap = "MaxHeap"
	minHeap = "MinHeap"
)

type Heap[T any] struct {
	list    *arraylist.List[T]
	compare func(a, b T) int
	minHeap bool
}

func MaxHeap[T any](comp comparator.Comparator[T], vs ...T) *Heap[T] {
	h := &Heap[T]{
		list:    arraylist.New(vs...),
		compare: comp,
		minHeap: false,
	}

	h.buildHeap()

	return h
}

func MinHeap[T any](comp comparator.Comparator[T], vs ...T) *Heap[T] {
	h := &Heap[T]{
		list:    arraylist.New(vs...),
		compare: comp,
		minHeap: true,
	}

	h.buildHeap()

	return h
}

func (h *Heap[T]) Name() string {
	switch h.minHeap {
	case true:
		return minHeap
	default:
		return maxHeap
	}
}

func (h *Heap[T]) Size() int {
	return h.list.Size()
}

func (h *Heap[T]) Empty() bool {
	return h.list.Empty()
}

func (h *Heap[T]) Values() []T {
	return h.list.Values()
}

func (h *Heap[T]) String() string {
	return fmt.Sprintf("%v", h.Values())
}

func (h *Heap[T]) Reset() {
	h.list.Reset()
}

func (h *Heap[T]) Add(vs ...T) {
	h.list.Add(vs...)
	h.buildHeap()
}

func (h *Heap[T]) Insert(vs ...T) {
	h.Add(vs...)
}

func (h *Heap[T]) Remove(i int) (T, bool) {
	var value T
	if h.Empty() || i < 0 || i >= h.Size() {
		return value, false
	}

	h.list.Swap(i, h.Size()-1)
	value, ok := h.list.Remove(h.Size() - 1)
	h.buildHeap()
	return value, ok
}

func (h *Heap[T]) Pop() (T, bool) {
	return h.Remove(0)
}

func (h *Heap[T]) buildHeap() {
	h.heapify()
	h.rightify()
}

func (h *Heap[T]) heapify() {
	for i := (h.Size() / 2) - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *Heap[T]) rightify() {
	for i := 0; i < (h.Size() / 2); i++ {
		child := 2*i + 1
		childValue, ok := h.list.Get(child)
		if ok {
			sibling := child + 1
			siblingValue, ok := h.list.Get(sibling)
			if ok && h.compare(siblingValue, childValue) == -1 {
				h.list.Swap(child, sibling)
				h.siftDown(child)
				h.siftDown(sibling)
			}
		}
	}
}

func (h *Heap[T]) siftDown(parent int) {
	child := 2*parent + 1
	childValue, ok := h.list.Get(child)

	if ok {
		sibling := child + 1
		siblingValue, ok := h.list.Get(sibling)

		if ok && h.minHeap && h.compare(siblingValue, childValue) == -1 || !h.minHeap && h.compare(siblingValue, childValue) == 1 {
			child = sibling
			childValue = siblingValue
		}
		parentValue, ok := h.list.Get(parent)
		if ok && h.minHeap && h.compare(parentValue, childValue) == 1 || !h.minHeap && h.compare(parentValue, childValue) == -1 {
			h.list.Swap(parent, child)
			h.siftDown(child)
		}
	}
}
