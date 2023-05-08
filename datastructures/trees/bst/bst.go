package bst

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/queues/linkedlistqueue"
	"github.com/mhrdini/godsa/datastructures/utils/comparator"
)

const bst = "BST"

type Tree[T any] struct {
	size    int
	root    *Node[T]
	compare func(a, b T) int
}

type Node[T any] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

type traversal[T any] func(n *Node[T], ch chan T)

func New[T any](comp comparator.Comparator[T], vs ...T) *Tree[T] {
	t := &Tree[T]{size: 0, root: nil, compare: comp}
	for _, v := range vs {
		t.Insert(v)
	}
	return t
}

func (t *Tree[T]) Name() string {
	return bst
}

func (t *Tree[T]) Size() int {
	return t.size
}

func (t *Tree[T]) Empty() bool {
	return t.size == 0 && t.root == nil
}

func (t *Tree[T]) Values() []T {
	return t.Traverse(InOrder[T])
}

func (t *Tree[T]) Traverse(walk traversal[T]) []T {
	vs := []T{}
	ch := make(chan T)
	go walk(t.root, ch)
	for i := 0; i < t.size; i++ {
		v := <-ch
		vs = append(vs, v)
	}
	return vs
}

func (t *Tree[T]) String() string {
	return fmt.Sprintf("%v", t.Values())
}

func (t *Tree[T]) Reset() {
	t.size = 0
	t.root = nil
}

func (t *Tree[T]) Insert(v T) {
	newNode := &Node[T]{value: v, left: nil, right: nil}

	if t.Empty() {
		t.root = newNode
	} else {
		t.root.Insert(newNode, t.compare)
	}
	t.size++
}

func (n *Node[T]) Insert(newNode *Node[T], compare comparator.Comparator[T]) {
	switch result := compare(newNode.value, n.value); result {
	case 1:
		if n.right == nil {
			n.right = newNode
		} else {
			n.right.Insert(newNode, compare)
		}
	default:
		if n.left == nil {
			n.left = newNode
		} else {
			n.left.Insert(newNode, compare)
		}
	}
}
func InOrder[T any](n *Node[T], ch chan T) {
	if n.left != nil {
		InOrder(n.left, ch)
	}
	ch <- n.value
	if n.right != nil {
		InOrder(n.right, ch)
	}
}

func PreOrder[T any](n *Node[T], ch chan T) {
	ch <- n.value
	if n.left != nil {
		PreOrder(n.left, ch)
	}
	if n.right != nil {
		PreOrder(n.right, ch)
	}
}

func PostOrder[T any](n *Node[T], ch chan T) {
	if n.left != nil {
		PostOrder(n.left, ch)
	}
	if n.right != nil {
		PostOrder(n.right, ch)
	}
	ch <- n.value
}

func LevelOrder[T any](n *Node[T], ch chan T) {
	queue := linkedlistqueue.New[*Node[T]]()
	queue.Enqueue(n)
	for !queue.Empty() {
		v, _ := queue.Dequeue()
		ch <- v.value
		if v.left != nil {
			queue.Enqueue(v.left)
		}
		if v.right != nil {
			queue.Enqueue(v.right)
		}
	}
}
