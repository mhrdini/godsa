package linkedlistqueue

import (
	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
)

const linkedListQueue = "LinkedListQueue"

type Queue[T any] struct {
	list *singlylinkedlist.List[T]
}

func New[T any](vs ...T) *Queue[T] {
	q := &Queue[T]{list: &singlylinkedlist.List[T]{}}
	for _, v := range vs {
		q.Enqueue(v)
	}
	return q
}

func (q *Queue[T]) Name() string {
	return linkedListQueue
}

func (q *Queue[T]) Size() int {
	return q.list.Size()
}

func (q *Queue[T]) Empty() bool {
	return q.list.Empty()
}

func (q *Queue[T]) Values() []T {
	return q.list.Values()
}

func (q *Queue[T]) String() string {
	return q.list.String()
}

func (q *Queue[T]) Reset() {
	q.list.Reset()
}

func (q *Queue[T]) Enqueue(v T) {
	q.list.Add(v)
}

func (q *Queue[T]) Dequeue() (v T, ok bool) {
	v, ok = q.list.Get(0)
	if ok {
		q.list.Remove(0)
	}
	return
}

func (q *Queue[T]) Peek() (v T, ok bool) {
	return q.list.Get(0)
}
