package fibonacciheap

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/trees"
	"github.com/mhrdini/godsa/datastructures/utils/comparator"
)

/* --------------------------------------------------------------------------
 * A fibonacci heap is a forest of fibonacci trees.
 *
 * It consists of a root list: an arbitrarily-ordered, circular, doubly-linked
 * list of root nodes, where each root node may or may not be the root of a
 * fibonacci tree.
 *
 * Each node may have children arranged in a child list: an arbitrarily-ordered,
 * circular, doubly-linked list of child nodes.
 *
 * A fibonacci heap is accessed via the heap's min/max node, i.e. the root node of
 * the tree containing the min/max value. If more than one root shares the same
 * min/max value, then any such root may serve as the min/max node.
/* -------------------------------------------------------------------------- */

const (
	fibonacci = "Fibonacci"
	maxHeap   = "MaxHeap"
	minHeap   = "MinHeap"
)

type FibonacciHeap[T any] struct {
	minOrMaxRoot *FibonacciNode[T]
	size         int // total # of nodes
	trees        int // total # of rooted trees
	marked       int // total # of marked nodes
	minHeap      bool
	compare      comparator.Comparator[T]
}

type FibonacciNode[T any] struct {
	value  T
	degree int // Total number of children

	// A node x is marked if it has lost a child since the last time x was
	// made the child of another node.
	// - Newly created nodes are unmarked.
	// - A node x becomes unmarked whenever it is made the child of another node
	marked bool

	parent *FibonacciNode[T]
	child  *FibonacciNode[T] // Points to any one of its children

	// The children of each node x are linked together in a circular, doubly
	// linked list, called a child list:
	// - Each child y has pointers y.left and y.right that point to y's left and right siblings.
	// - Siblings may appear in a child list in any order.
	// - If a child node y is an only child, then y.left = y.right = y
	left  *FibonacciNode[T]
	right *FibonacciNode[T]
}

func MinHeap[T any](compare comparator.Comparator[T]) *FibonacciHeap[T] {
	h := &FibonacciHeap[T]{
		minOrMaxRoot: nil,
		size:         0,
		trees:        0,
		marked:       0,
		minHeap:      true,
		compare:      compare,
	}
	return h
}

func MaxHeap[T any](compare comparator.Comparator[T]) *FibonacciHeap[T] {
	h := &FibonacciHeap[T]{
		minOrMaxRoot: nil,
		size:         0,
		trees:        0,
		marked:       0,
		minHeap:      false,
		compare:      compare,
	}
	return h
}

func NewEmptyHeap[T any](minHeap bool, compare comparator.Comparator[T]) *FibonacciHeap[T] {
	if minHeap {
		return MinHeap(compare)
	} else {
		return MaxHeap(compare)
	}
}

func NewNode[T any](v T) *FibonacciNode[T] {
	n := &FibonacciNode[T]{
		value:  v,
		degree: 0,
		marked: false,
		parent: nil,
		child:  nil,
	}
	n.left = n
	n.right = n
	return n
}

/* -------------------------------------------------------------------------- */
/*                               HEAP INSPECTION                              */
/* -------------------------------------------------------------------------- */

func (h *FibonacciHeap[T]) Name() string {
	var minMax string
	if h.minHeap {
		minMax = minHeap
	} else {
		minMax = maxHeap
	}
	return fmt.Sprintf("%v%v", fibonacci, minMax)
}

func (h *FibonacciHeap[T]) Size() int {
	return h.size
}

/* -------------------------------------------------------------------------- */
/*                               NODE INSPECTION                              */
/* -------------------------------------------------------------------------- */

func (n *FibonacciNode[T]) Value() (value T, ok bool) {
	if n != nil {
		value = n.value
		ok = true
	}
	return
}

func (n *FibonacciNode[T]) Children() []trees.INode[T] {
	return []trees.INode[T]{
		0: n.child,
		1: n.left,
		2: n.right,
	}
}

func (n *FibonacciNode[T]) IsNil() bool {
	return n == nil
}

func (n *FibonacciNode[T]) String() string {
	return fmt.Sprintf("(%v, %v)", n.value, n.degree)
}

/* -------------------------------------------------------------------------- */
/*                              INSERTION/REMOVAL                             */
/* -------------------------------------------------------------------------- */
