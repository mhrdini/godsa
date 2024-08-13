package binomialheap

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/trees"
	"github.com/mhrdini/godsa/datastructures/utils/comparator"
)

// https://brilliant.org/wiki/binomial-heap/

/* --------------------------------------------------------------------------
 * A binomial heap is a forest of binomial trees.
 *
 * A binomial tree of degree 0 is a single node.
 * A binomial tree of degree k has a root node with k children. The degree of
 * those children are k-1, k-2,..., 2, 1, 0.
 *
 * A binomial tree of degree k has:
 * - 2^k nodes
 * - a height of k
 * - kCi nodes at depth i for i = [0, 1]
 * - a root of degree k, greater than any other node
 *
 * A binomial heap is a forest of binomial trees that satisfy the heap invariant.
 * There can be only 0 or 1 binomial tree of degree k in the forest.
 *
 * The two key features of binomial heaps are:
 *
 * 1. The roots of the forest are <= log(n)
 * 2. Merging to heaps is binary addition
 *
 * This brings merge() from O(n + m) in binary heaps to O(log n + log m) in
 * binomial heaps.
 *
 * enqueue() -> O(log n)
 * extractMin() -> O(log n)
 * findMin() -> O(1)
 * merge -> O(log n + log m)
 *
 * Implementation based on Binomial Heap pseudocode from CLRS.
/* -------------------------------------------------------------------------- */

const (
	binomial = "Binomial"
	maxHeap  = "MaxHeap"
	minHeap  = "MinHeap"
)

type BinomialHeap[T any] struct {
	head         *BinomialNode[T]
	size         int
	minHeap      bool
	minOrMaxRoot *BinomialNode[T]

	// smallest or largest value for deleteNode(node)
	// deleteNode will decrease the node to the smallest/largest value so it swims
	// up to the root, and then calls dequeue()
	minOrMaxValue T

	compare comparator.Comparator[T]
}

type BinomialNode[T any] struct {
	value   T
	degree  int // no of children
	parent  *BinomialNode[T]
	child   *BinomialNode[T]
	sibling *BinomialNode[T]
}

func MaxHeap[T any](comp comparator.Comparator[T], minOrMaxValue T) *BinomialHeap[T] {
	h := &BinomialHeap[T]{
		head:          nil,
		minOrMaxRoot:  nil,
		size:          0,
		minHeap:       false,
		minOrMaxValue: minOrMaxValue,
		compare:       comp,
	}

	return h
}

func MinHeap[T any](comp comparator.Comparator[T], minOrMaxValue T) *BinomialHeap[T] {
	h := &BinomialHeap[T]{
		head:          nil,
		minOrMaxRoot:  nil,
		size:          0,
		minHeap:       true,
		minOrMaxValue: minOrMaxValue,
		compare:       comp,
	}

	return h
}

func NewEmptyHeap[T any](minHeap bool, comp comparator.Comparator[T], minOrMaxValue T) *BinomialHeap[T] {
	if minHeap {
		return MinHeap(comp, minOrMaxValue)
	} else {
		return MaxHeap(comp, minOrMaxValue)
	}
}

func NewNode[T any](v T) *BinomialNode[T] {
	n := &BinomialNode[T]{
		value:   v,
		degree:  0,
		parent:  nil,
		child:   nil,
		sibling: nil,
	}
	return n
}

/* -------------------------------------------------------------------------- */
/*                               HEAP INSPECTION                              */
/* -------------------------------------------------------------------------- */

func (h *BinomialHeap[T]) Name() string {
	var minMax string
	if h.minHeap {
		minMax = minHeap
	} else {
		minMax = maxHeap
	}
	return fmt.Sprintf("%v%v", binomial, minMax)
}

func (h *BinomialHeap[T]) Size() int {
	return h.size
}

func (h *BinomialHeap[T]) Empty() bool {
	return h.size == 0
}

func (h *BinomialHeap[T]) Values() []T {
	return trees.Traverse(trees.ITree[T](h), trees.LevelOrder[T])
}

func (h *BinomialHeap[T]) String() string {
	return fmt.Sprintf("%v", h.Values())
}

func (h *BinomialHeap[T]) Reset() {
	h.head = nil
	h.minOrMaxRoot = nil
	h.size = 0
}

func (h *BinomialHeap[T]) Root() trees.INode[T] {
	return h.head
}

/* -------------------------------------------------------------------------- */
/*                               NODE INSPECTION                              */
/* -------------------------------------------------------------------------- */

func (n *BinomialNode[T]) Value() (value T, ok bool) {
	if n != nil {
		value = n.value
		ok = true
	}
	return
}

func (n *BinomialNode[T]) Children() []trees.INode[T] {
	return []trees.INode[T]{
		0: n.child,
		1: n.sibling,
	}
}

func (n *BinomialNode[T]) IsNil() bool {
	return n == nil
}

func (n *BinomialNode[T]) String() string {
	return fmt.Sprintf("(%v, %v)", n.value, n.degree)
}

/* -------------------------------------------------------------------------- */
/*                             INSERTION/DELETION                             */
/* -------------------------------------------------------------------------- */

func (h *BinomialHeap[T]) Insert(v T) {

	var hPrime *BinomialHeap[T]
	var vComparedToSmallestValue = h.compare(v, h.minOrMaxValue)

	switch h.minHeap {
	case true:
		if vComparedToSmallestValue == comparator.Lesser {
			hPrime = MinHeap(h.compare, v)
		} else {
			hPrime = MinHeap(h.compare, h.minOrMaxValue)
		}
	case false:
		if vComparedToSmallestValue == comparator.Greater {
			hPrime = MaxHeap(h.compare, v)
		} else {
			hPrime = MaxHeap(h.compare, h.minOrMaxValue)
		}
	}

	hPrime.head = NewNode(v)
	hPrime.size++

	//	fmt.Printf("/* -------------------------------------------------------------------------- */\n\n")
	//	fmt.Printf("Inserting %v ...\n\n", v)
	h.head, h.size, h.minOrMaxRoot, h.minOrMaxValue = union(h, hPrime)
	// fmt.Printf("%v\n\n", h)
}

func (h *BinomialHeap[T]) Remove(v T) {
}

/* -------------------------------------------------------------------------- */
/*                                    UNION                                   */
/* -------------------------------------------------------------------------- */

// Merges two forests and returns one forest monotonically sorted by degree
// in O(t) where t is the total number of trees in both forests.
func merge[T any](h1, h2 *BinomialHeap[T]) *BinomialNode[T] {
	if h1 == nil || h1.head == nil {
		return h2.head
	}

	if h2 == nil || h2.head == nil {
		return h1.head
	}

	var head *BinomialNode[T]
	a := h1.head
	b := h2.head

	if a.degree < b.degree {
		head = a
		a = a.sibling
	} else {
		head = b
		b = b.sibling
	}

	no := 0
	curr := head
	//	fmt.Printf("%v\tcurr: %v\ta: %v\tb: %v\n", no, curr, a, b)

	for a != nil && b != nil {
		no++
		if a.degree < b.degree {
			curr.sibling = a
			a = a.sibling
		} else {
			curr.sibling = b
			b = b.sibling
		}
		curr = curr.sibling
		//	fmt.Printf("%v\tcurr: %v\ta: %v\tb: %v\n", no, curr, a, b)
	}

	if a != nil {
		curr.sibling = a
	} else {
		curr.sibling = b
	}

	//	fmt.Printf("END\thead: %v\ta: %v\tb: %v\n\n", head, a, b)

	return head
}

// func merge[T any](h1, h2 *BinomialHeap[T]) *BinomialNode[T] {
// 	var root *BinomialNode[T] = nil
// 	var pos **BinomialNode[T] = &root

// 	a := h1.head
// 	b := h2.head

// 	no := 0
// //	fmt.Printf("%v\t*pos: %v\ta: %v\tb: %v\n", no, *pos, a, b)

// 	for a != nil && b != nil {
// 		if a.degree < b.degree {
// 			*pos = a
// 			a = a.sibling
// 		} else {
// 			*pos = b
// 			b = b.sibling
// 		}

// 		no++
// 	//	fmt.Printf("%v\t*pos: %v\ta: %v\tb: %v\n", no, *pos, a, b)
// 	}

// 	if a != nil {
// 		*pos = a
// 	} else {
// 		*pos = b
// 	}

// //	fmt.Printf("END\t*pos: %v\ta: %v\tb: %v\n\n", *pos, a, b)

// 	return root
// }

// Links a B_(k-1) tree rooted at node y to the B_k-1 tree rooted at node z, i.e.
// make z the parent of y. Thus, z becomes the root of a B_k tree.
func link[T any](y, z *BinomialNode[T]) {
	y.parent = z
	y.sibling = z.child
	z.child = y
	z.degree++
}

// Unites two heaps, h1 and h2, in two phases:
/* ------------------------------- First Phase ------------------------------ */
//
// - Performed by call to merge(h1, h2) -> O(log n)
// - Merges the root lists of h1 and h2 into a single heap h of trees sorted
//  by degree in a monotonically increasing order
// - At this point there may be as many but no more than two roots of each
// degree that remains
//
/* ------------------------------ Second Phase ------------------------------ */
//
// - Using a loop, we go through the list of roots and coalesce the trees such that
// there will exist only one tree of each degree in the heap.
// - Initially, we start at the leftmost root, keeping it as 'curr', and its
// sibling, keeping it as 'next'.
// - Invariant of the loop: Both curr and next being non-nil.
//
// There exists 4 cases that may occur at each iteration:
//
// == Case 1: curr.degree != next.degree
// - No linking between curr and next
// - We move pointers one further down the list
//
// == Case 2: curr.degree == next.degree == next.sibling.degree, i.e. when curr
// is the first of 3 roots of equal degree
// - No linking between curr and next
// - We move pointers one further down the list
//
// == Case 3: curr.degree == next.degree != next.sibling.degree AND
// curr.value (<= for min heap, > for max heap) next.value
// - next is linked to curr -> curr is made the root of next
// - next is removed from the root list
// - next is made the leftmost child of curr
//
// == Case 4: curr.degree == next.degree != next.sibling.degree AND
// curr.value (> for min heap, <= for max heap) next.value
// - curr is linked to next -> next is made the root of curr, next's
// - curr is removed from the root list
// - curr is made the leftmost child of next
// - We move the curr pointer one further down the list
//
func union[T any](h1, h2 *BinomialHeap[T]) (head *BinomialNode[T],
	size int,
	minOrMaxRoot *BinomialNode[T],
	minOrMaxValue T) {

	//	fmt.Printf("union PRE\th1: %v\th2: %v\n\n", h1.head, h2.head)

	newMinOrMaxValue := getMinOrMax(h1.minHeap, h1.compare, h1.minOrMaxValue, h2.minOrMaxValue)
	h := NewEmptyHeap(h1.minHeap, h1.compare, newMinOrMaxValue)
	// first phase
	h.size = h1.size + h2.size
	h.head = merge(h1, h2)

	if h.head == nil {
		return h.head, h.size, h.minOrMaxRoot, h.minOrMaxValue
	}

	var prev, curr, next *BinomialNode[T]
	no := 0
	prev = nil
	curr = h.head
	next = curr.sibling

	//	fmt.Printf("%v -> PRE\n> prev: %v\n> curr: %v\n> next: %v\n\n", no, prev, curr, next)
	for next != nil {
		no++
		// cases 1 + 2
		if curr.degree != next.degree || next.sibling != nil && next.sibling.degree == curr.degree {
			prev = curr
			curr = next
			//	fmt.Printf("%v -> CASES 1 + 2\n> prev: %v\n> curr: %v\n> next: %v\n\n", no, prev, curr, next)
			// case 3 + 4
		} else {
			currComparedToNext := h.compare(curr.value, next.value)
			//  case 3
			if (h.minHeap && currComparedToNext != comparator.Greater) || (!h.minHeap && currComparedToNext != comparator.Lesser) {
				//	fmt.Printf("%v -> PRE LINK CASE 3\n> prev: %v\n> curr: %v\n> next: %v\n\n", no, prev, curr, next)
				curr.sibling = next.sibling
				link(next, curr)
				//	fmt.Printf("%v -> POST LINK CASE 3\n> prev: %v\n> curr: %v\n> next: %v\n\n", no, prev, curr, next)
				// case 4
			} else {
				if prev == nil {
					h.head = next
				} else {
					prev.sibling = next
				}
				link(curr, next)
				curr = next
				//	fmt.Printf("%v -> CASE 4\n> prev: %v\n> curr: %v\n> next: %v\n\n", no, prev, curr, next)
			}
		}
		next = curr.sibling
		//	fmt.Printf("%v\n\n", h)
	}
	head = h.head
	size = h.size
	minOrMaxRoot = h.minOrMaxRoot
	minOrMaxValue = h.minOrMaxValue
	return
}

/* -------------------------------------------------------------------------- */
/*                              HELPER FUNCTIONS                              */
/* -------------------------------------------------------------------------- */

func getMinOrMax[T any](minHeap bool, comp comparator.Comparator[T], x, y T) T {
	result := comp(x, y)
	if minHeap && result == comparator.Lesser || !minHeap && result == comparator.Greater {
		return x
	} else {
		return y
	}
}
