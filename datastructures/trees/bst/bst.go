package bst

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/trees"
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
	return trees.Traverse(trees.ITree[T](t), trees.InOrder[T])
}

func (t *Tree[T]) String() string {
	return fmt.Sprintf("%v", t.Values())
}

func (t *Tree[T]) Reset() {
	t.size = 0
	t.root = nil
}

func (t *Tree[T]) Root() trees.INode[T] {
	return t.root
}

func (t *Tree[T]) Insert(v T) {
	t.root = t.root.insert(t, v)
}

func (n *Node[T]) Value() (value T, ok bool) {
	if n == nil {
		return
	}
	value = n.value
	ok = true
	return
}

func (n *Node[T]) Left() trees.INode[T] {
	if n == nil {
		return nil
	}
	return n.left
}

func (n *Node[T]) Right() trees.INode[T] {
	if n == nil {
		return nil
	}
	return n.right
}

func (n *Node[T]) insert(tree *Tree[T], v T) *Node[T] {
	if n == nil {
		tree.size++
		return &Node[T]{value: v, left: nil, right: nil}
	}

	switch result := tree.compare(v, n.value); result {
	case -1:
		n.left = n.left.insert(tree, v)
	case 1:
		n.right = n.right.insert(tree, v)
	}
	return n
}
