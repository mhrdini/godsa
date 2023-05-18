package redblacktree

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/trees"
	"github.com/mhrdini/godsa/datastructures/utils/comparator"
)

const redblacktree = "RedBlackTree"

type color bool

const (
	red   = color(true)
	black = color(false)
)

type Tree[T any] struct {
	size    int
	root    *Node[T]
	null    *Node[T]
	compare comparator.Comparator[T]
}

type Node[T any] struct {
	value  T
	color  color
	left   *Node[T]
	right  *Node[T]
	parent *Node[T]
}

func New[T any](comp comparator.Comparator[T], vs ...T) trees.ITree[T] {
	var zeroValue T
	null := &Node[T]{zeroValue, black, nil, nil, nil}
	t := &Tree[T]{0, null, null, comp}
	for _, v := range vs {
		t.Insert(v)
	}

	return t
}

func (t *Tree[T]) Name() string {
	return redblacktree
}

func (t *Tree[T]) Size() int {
	return t.size
}

func (t *Tree[T]) Empty() bool {
	return t.size == 0 && t.root == t.null
}

func (t *Tree[T]) Values() []T {
	return trees.Traverse(trees.ITree[T](t), trees.InOrder[T])
}

func (t *Tree[T]) String() string {
	return fmt.Sprintf("%v", t.Values())
}

func (t *Tree[T]) Reset() {
	t.size = 0
	t.root = t.null
}

func (t *Tree[T]) Root() trees.INode[T] {
	return t.root
}

func (n *Node[T]) Value() (value T, ok bool) {
	if n.IsNil() {
		return
	}
	value = n.value
	ok = true
	return
}

func (n *Node[T]) Children() []trees.INode[T] {
	return []trees.INode[T]{
		0: n.left,
		1: n.right,
	}
}

func (n *Node[T]) IsNil() bool {
	return n.color == black && n.left == nil && n.right == nil && n.parent == nil
}

func (t *Tree[T]) Search(n *Node[T], v T) (*Node[T], bool) {
	node := n
	for node != t.null {
		switch result := t.compare(v, node.value); result {
		case -1:
			return t.Search(node.left, v)
		case 1:
			return t.Search(node.right, v)
		default:
			return node, true
		}
	}
	return t.null, false
}

func (t *Tree[T]) Insert(v T) {
	newNode := &Node[T]{v, red, t.null, t.null, t.null}

	currentNode := t.root
	parent := t.null // parent will be the parent of z

	for currentNode != t.null {
		parent = currentNode
		switch t.compare(newNode.value, currentNode.value) {
		case -1:
			currentNode = currentNode.left
		default:
			currentNode = currentNode.right
		}
	}

	t.size++

	newNode.parent = parent

	if parent == t.null {
		t.root = newNode
	} else if t.compare(newNode.value, parent.value) == -1 {
		parent.left = newNode
	} else {
		parent.right = newNode
	}

	if newNode.parent == t.null {
		newNode.color = black
		return
	}

	if newNode.parent.parent == t.null {
		return
	}

	t.fixPostInsert(newNode)
}

func (t *Tree[T]) Remove(v T) {

	var replacementSuccessor = t.null
	removed, ok := t.Search(t.root, v)
	if ok {
		t.size--
		replacement := removed
		replacementColor := replacement.color
		if removed.left == t.null {
			replacementSuccessor = removed.right
			t.transplant(removed, removed.right)
		} else if removed.right == t.null {
			replacementSuccessor = removed.left
			t.transplant(removed, removed.right)
		} else {
			replacement = t.minimum(removed.right)
			replacementColor = replacement.color
			replacementSuccessor = replacement.right
			if replacement.parent == removed {
				replacementSuccessor.parent = replacement
			} else {
				t.transplant(replacement, replacement.right)
				replacement.right = removed.left
				replacement.right.parent = replacement
			}
			t.transplant(removed, replacement)
			replacement.left = removed.left
			replacement.left.parent = replacement
			replacement.color = removed.color
		}
		if replacementColor == black {
			t.fixPostRemove(replacementSuccessor)
		}
	}
}

func (t *Tree[T]) fixPostInsert(n *Node[T]) {

	node := n
	var uncle = t.null

	for node.parent.color == red {
		switch node.parent {
		case node.parent.parent.left:
			uncle = node.parent.parent.right
			switch uncle.color {
			case red:
				// case 1
				node.parent.color = black
				uncle.color = black
				node.parent.parent.color = red
				node = node.parent.parent
			case black:
				// case 2
				if node == node.parent.right {
					node = node.parent
					t.leftRotate(node)
				}
				// case 3
				node.parent.color = black
				node.parent.parent.color = red
				t.rightRotate(node.parent.parent)
			}
		case node.parent.parent.right:
			uncle = node.parent.parent.left
			switch uncle.color {
			case red:
				// case 1
				node.parent.color = black
				uncle.color = black
				node.parent.parent.color = red
				node = node.parent.parent
			case black:
				// case 2
				if node == node.parent.left {
					node = node.parent
					t.rightRotate(n)
				}
				// case 3
				node.parent.color = black
				node.parent.parent.color = red
				t.leftRotate(node.parent.parent)
			}
		}
	}
	t.root.color = black
}

func (t *Tree[T]) fixPostRemove(n *Node[T]) {

	var sibling = t.null
	node := n

	for node != t.root && node.color == black {
		if node.parent != t.null {
			switch node {
			case node.parent.left:
				sibling = node.parent.right
				// case 1
				if sibling.color == red {
					sibling.color = black
					node.parent.color = red
					t.leftRotate(node.parent)
					sibling = node.parent.right
				}
				// case 2
				if sibling.left.color == black && sibling.right.color == black {
					sibling.color = red
					node = node.parent
				} else {
					// case 3
					if sibling.right.color == black {
						sibling.left.color = black
						sibling.color = red
						t.rightRotate(sibling)
						sibling = node.parent.right
					}
					// case 4
					sibling.color = node.parent.color
					node.parent.color = black
					sibling.right.color = black
					t.leftRotate(node.parent)
					node = t.root
				}
			case node.parent.right:
				sibling = node.parent.left
				// case 1
				if sibling.color == red {
					sibling.color = black
					node.parent.color = red
					t.rightRotate(node.parent)
					sibling = node.parent.left
				}
				// case 2
				if sibling.right.color == black && sibling.left.color == black {
					sibling.color = red
					node = node.parent
				} else {
					// case 3
					if sibling.left.color == black {
						sibling.right.color = black
						sibling.color = red
						t.leftRotate(sibling)
						sibling = node.parent.left
					}
					// case 4
					sibling.color = node.parent.color
					node.parent.color = black
					sibling.left.color = black
					t.rightRotate(node.parent)
					node = t.root
				}
			}
		}
	}
	node.color = black
}

func (t *Tree[T]) leftRotate(n *Node[T]) {
	newRoot := n.right
	n.right = newRoot.left
	if newRoot.left != t.null {
		newRoot.left.parent = n
	}
	newRoot.parent = n.parent
	if n.parent == t.null {
		t.root = newRoot
	} else if n == n.parent.left {
		n.parent.left = newRoot
	} else {
		n.parent.right = newRoot
	}
	newRoot.left = n
	n.parent = newRoot
}

func (t *Tree[T]) rightRotate(n *Node[T]) {

	newRoot := n.left
	n.left = newRoot.right
	if newRoot.right != t.null {
		newRoot.right.parent = n
	}
	newRoot.parent = n.parent
	if n.parent == t.null {
		t.root = newRoot
	} else if n == n.parent.right {
		n.parent.right = newRoot
	} else {
		n.parent.left = newRoot
	}
	newRoot.right = n
	n.parent = newRoot
}

func (t *Tree[T]) transplant(original, replacement *Node[T]) {
	if original.parent == t.null {
		t.root = replacement
	} else if original == replacement.parent.left {
		original.parent.left = replacement
	} else {
		original.parent.right = replacement
	}
	replacement.parent = original.parent
}

func (t *Tree[T]) minimum(n *Node[T]) *Node[T] {
	node := n
	for node.left != t.null {
		node = node.left
	}
	return node
}
