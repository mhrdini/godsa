package avltree

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/trees"
	"github.com/mhrdini/godsa/datastructures/utils/comparator"
)

const avltree = "AVLTree"

type Tree[T any] struct {
	size    int
	root    *Node[T]
	compare func(a, b T) int
}

type Node[T any] struct {
	value  T
	height int
	left   *Node[T]
	right  *Node[T]
}

func New[T any](comp comparator.Comparator[T], vs ...T) *Tree[T] {
	t := &Tree[T]{size: 0, root: nil, compare: comp}
	if len(vs) > 0 {
		for _, v := range vs {
			t.Insert(v)
		}
	}
	return t
}

func (t *Tree[T]) Name() string {
	return avltree
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

func (t *Tree[T]) Insert(v T) {
	t.root = t.root.insert(t, v)
}

func (n *Node[T]) insert(tree *Tree[T], v T) *Node[T] {
	if n == nil {
		tree.size++
		return &Node[T]{value: v, height: 1, left: nil, right: nil}
	}

	switch result := tree.compare(v, n.value); result {
	case -1:
		n.left = n.left.insert(tree, v)
	case 1:
		n.right = n.right.insert(tree, v)
	}
	return n.balance()
}

func (n *Node[T]) balance() *Node[T] {
	if n == nil {
		return n
	}

	n.updateHeight()

	switch balanceFactor := n.balanceFactor(); balanceFactor {
	case 2:
		if n.left != nil && n.left.balanceFactor() == -1 {
			n.left = n.left.leftRotate()
		}
		return n.rightRotate()
	case -2:
		if n.right != nil && n.right.balanceFactor() == 1 {
			n.right = n.right.rightRotate()
		}
		return n.leftRotate()
	}
	return n
}

func (n *Node[T]) leftRotate() *Node[T] {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.updateHeight()
	newRoot.updateHeight()
	return newRoot
}

func (n *Node[T]) rightRotate() *Node[T] {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.updateHeight()
	newRoot.updateHeight()
	return newRoot
}

func (n *Node[T]) balanceFactor() int {
	return n.left.getHeight() - n.right.getHeight()
}

func (n *Node[T]) updateHeight() {
	n.height = maxOf(n.left.getHeight(), n.right.getHeight()) + 1
}

func (n *Node[T]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func maxOf(a, b int) int {
	if a > b {
		return a
	}
	return b
}
