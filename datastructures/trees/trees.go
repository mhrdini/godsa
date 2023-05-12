package trees

import (
	"github.com/mhrdini/godsa/datastructures/containers"
	"github.com/mhrdini/godsa/datastructures/queues/linkedlistqueue"
)

type ITree[T any] interface {
	containers.Container[T]
	Root() INode[T]
	Insert(v T)
	Remove(v T)
}

type INode[T any] interface {
	Value() (value T, ok bool)
	Children() []INode[T]
	Height() int
	// String() string
}

type Traverser[T any] func(n INode[T], ch chan T)

func InOrder[T any](n INode[T], ch chan T) {
	if value, ok := n.Value(); ok {
		children := n.Children()
		totalChildren := len(children)
		for i := 0; i < totalChildren-1; i++ {
			InOrder(children[i], ch)
		}
		ch <- value
		InOrder(children[totalChildren-1], ch)
	}
}

func PreOrder[T any](n INode[T], ch chan T) {
	if value, ok := n.Value(); ok {
		ch <- value
		children := n.Children()
		for _, child := range children {
			PreOrder(child, ch)
		}
	}
}

func PostOrder[T any](n INode[T], ch chan T) {
	if value, ok := n.Value(); ok {
		children := n.Children()
		for _, child := range children {
			PostOrder(child, ch)
		}
		ch <- value
	}
}

func LevelOrder[T any](n INode[T], ch chan T) {
	queue := linkedlistqueue.New[INode[T]]()
	queue.Enqueue(n)
	for !queue.Empty() {
		node, _ := queue.Dequeue()
		if value, ok := node.Value(); ok {
			ch <- value
			for _, child := range node.Children() {
				if child != nil {
					queue.Enqueue(child)
				}
			}
		}
	}
}

func Traverse[T any](t ITree[T], walk Traverser[T]) []T {
	vs := []T{}
	ch := make(chan T)
	go walk(t.Root(), ch)
	for i := 0; i < t.Size(); i++ {
		v := <-ch
		vs = append(vs, v)
	}
	return vs
}
