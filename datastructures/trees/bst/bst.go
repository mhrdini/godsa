package bst

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/utils"
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

func New[T any](comp utils.Comparator[T], vs ...T) *Tree[T] {
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
	return t.Traverse(Inorder[T])
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

func (n *Node[T]) Insert(newNode *Node[T], compare utils.Comparator[T]) {
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
func Inorder[T any](n *Node[T], ch chan T) {
	if n.left != nil {
		Inorder(n.left, ch)
	}
	ch <- n.value
	if n.right != nil {
		Inorder(n.right, ch)
	}
}

func Preorder[T any](n *Node[T], ch chan T) {
	ch <- n.value
	if n.left != nil {
		Preorder(n.left, ch)
	}
	if n.right != nil {
		Preorder(n.right, ch)
	}
}

func Postorder[T any](n *Node[T], ch chan T) {
	if n.left != nil {
		Postorder(n.left, ch)
	}
	if n.right != nil {
		Postorder(n.right, ch)
	}
	ch <- n.value
}
