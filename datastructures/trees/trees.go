package trees

import (
	"github.com/mhrdini/godsa/datastructures/containers"
	"github.com/mhrdini/godsa/datastructures/queues/linkedlistqueue"
)

type ITree[T any] interface {
	containers.Container[T]
	Root() INode[T]
}

type INode[T any] interface {
	Value() (value T, ok bool)
	Left() INode[T]
	Right() INode[T]
	// String() string
}

type Traverser[T any] func(n INode[T], ch chan T)

func InOrder[T any](n INode[T], ch chan T) {
	if n.Left() != nil {
		InOrder(n.Left(), ch)
	}
	if value, ok := n.Value(); ok {
		ch <- value
	}
	if n.Right() != nil {
		InOrder(n.Right(), ch)
	}
}

func PreOrder[T any](n INode[T], ch chan T) {
	if value, ok := n.Value(); ok {
		ch <- value
	}
	if n.Left() != nil {
		PreOrder(n.Left(), ch)
	}
	if n.Right() != nil {
		PreOrder(n.Right(), ch)
	}
}

func PostOrder[T any](n INode[T], ch chan T) {
	if n.Left() != nil {
		PostOrder(n.Left(), ch)
	}
	if n.Right() != nil {
		PostOrder(n.Right(), ch)
	}
	if value, ok := n.Value(); ok {
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
		}
		if node.Left() != nil {
			queue.Enqueue(node.Left())
		}
		if node.Right() != nil {
			queue.Enqueue(node.Right())
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
