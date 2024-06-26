package redblacktree

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/trees"
	"github.com/mhrdini/godsa/datastructures/utils/comparator"
)

const redBlackTree = "RedBlackTree"

type color int

const (
	red = color(iota)
	black
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
	return redBlackTree
}

func (t *Tree[T]) Size() int {
	return t.size
}

func (t *Tree[T]) Empty() bool {
	return t.size == 0 && t.root == t.null
}

func (t *Tree[T]) Values() []T {
	return trees.Traverse(trees.ITree[T](t), trees.LevelOrder[T])
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
		case comparator.Lesser:
			return t.Search(node.left, v)
		case comparator.Greater:
			return t.Search(node.right, v)
		default:
			return node, true
		}
	}
	return t.null, false
}

// Insert receives a value and creates a new red node out of it. It finds and inserts it
// into an initial position, and then uses the fixPostInsert function to correct the tree
// and fix violations against the RBT rules.
func (t *Tree[T]) Insert(v T) {
	newNode := &Node[T]{v, red, t.null, t.null, t.null}

	currentNode := t.root
	parent := t.null // parent will be the parent of newNode

	for currentNode != t.null {
		parent = currentNode
		switch t.compare(newNode.value, currentNode.value) {
		case comparator.Lesser:
			currentNode = currentNode.left
		default:
			currentNode = currentNode.right
		}
	}

	t.size++

	newNode.parent = parent

	if parent == t.null {
		t.root = newNode
	} else if t.compare(newNode.value, parent.value) == comparator.Lesser {
		parent.left = newNode
	} else {
		parent.right = newNode
	}

	// case 0: new node is the root, so color it black
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

	var replacement = t.null
	removed, ok := t.Search(t.root, v)

	if ok {
		t.size--

		successor := removed
		successorColor := successor.color

		// case 1: if left subtree of removed node is nil
		if removed.left == t.null {
			replacement = removed.right
			t.transplant(removed, replacement)
			// case 2: if right subtree of removed node is nil
		} else if removed.right == t.null {
			replacement = removed.left
			t.transplant(removed, replacement)
			// case 3: if neither subtree of removed node is nil
		} else {
			successor = t.minimum(removed.right)
			successorColor = successor.color
			replacement = successor.right

			if successor != removed.right {
				t.transplant(successor, successor.right)
				successor.right = removed.right
				successor.right.parent = successor
			} else {
				replacement.parent = successor
			}

			t.transplant(removed, successor)
			successor.left = removed.left
			successor.left.parent = successor
			successor.color = removed.color
		}
		if successorColor == black {
			t.fixPostRemove(replacement)
		}
	}
}

// fixPostInsert fixes violations in the RBT after the insertion of a node N by identifying 3 cases:
// - Case 1 -> Where the uncle node of N is red
// - Case 2 -> Where the uncle node of N is black, and either N is a left child and its parent is a
// right child or vice versa (forming a triangle)
// - Case 3 -> Where the uncle node of N is black, and either both N and its parent are right
// children or both are left children (forming a line)
// The fix for each case is as follows:
// - Case 1 -> Recolour N's parent and uncle black, and recolour N's grandparent red
// - Case 2 -> Rotate N's parent in the opposite direction of N so that N takes the place of its parent
// - Case 3 -> Rotate N's grandparent in the opposite direction of N so that N's parent takes the
// place of its grandparent, then recolor N's original parent and grandparent
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

	for n != t.root && n.color == black {
		if n.parent != t.null {

			parent := n.parent
			n.parent = nil

			switch n {
			case parent.left:
				sibling := parent.right
				// case 1
				if sibling.color == red {
					sibling.color = black
					parent.color = red
					t.leftRotate(parent)
					sibling = parent.right
				}
				// case 2
				if sibling.left.color == black && sibling.right.color == black {
					sibling.color = red
					n = parent
				} else {
					// case 3
					if sibling.right.color == black {
						sibling.left.color = black
						sibling.color = red
						t.rightRotate(sibling)
						sibling = parent.right
					}
					// case 4
					sibling.color = parent.color
					parent.color = black
					sibling.right.color = black
					t.leftRotate(parent)
					n = t.root
				}
			case parent.right:
				sibling := parent.left
				// case 1
				if sibling.color == red {
					sibling.color = black
					parent.color = red
					t.rightRotate(parent)
					sibling = parent.left
				}
				// case 2
				if sibling.right.color == black && sibling.left.color == black {
					sibling.color = red
					n = parent
				} else {
					// case 3
					if sibling.left.color == black {
						sibling.right.color = black
						sibling.color = red
						t.leftRotate(sibling)
						sibling = parent.left
					}
					// case 4
					sibling.color = parent.color
					parent.color = black
					sibling.left.color = black
					t.rightRotate(parent)
					n = t.root
				}
			}
		}
	}
	n.color = black
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

// transplant is called by the Remove function to move subtrees, specifically in replacing the original
// node with the replacement node, whereby the replacement node becomes the original node's parent's
// new child. Note that other updates with respect to the original node's other children (which may
// have been affected by the current transplant function) are the responsibility of the calling function.
func (t *Tree[T]) transplant(original, replacement *Node[T]) {

	if original.parent == t.null {
		t.root = replacement
	} else if original == original.parent.left {
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
