package linkedliststack

import (
	"github.com/mhrdini/godsa/datastructures/lists/singlylinkedlist"
)

const linkedListStack = "LinkedListStack"

type Stack[T any] struct {
	list *singlylinkedlist.List[T]
}

func New[T any](vs ...T) *Stack[T] {
	s := &Stack[T]{list: &singlylinkedlist.List[T]{}}
	for _, v := range vs {
		s.Push(v)
	}
	return s
}

func (s *Stack[T]) Name() string {
	return linkedListStack
}

func (s *Stack[T]) Size() int {
	return s.list.Size()
}

func (s *Stack[T]) Empty() bool {
	return s.list.Empty()
}

func (s *Stack[T]) Values() []T {
	return s.list.Values()
}

func (s *Stack[T]) String() string {
	return s.list.String()
}

func (s *Stack[T]) Reset() {
	s.list.Reset()
}

func (s *Stack[T]) Push(v T) {
	s.list.Add(v)
}

func (s *Stack[T]) Pop() (v T, ok bool) {
	v, ok = s.list.Get(s.list.Size() - 1)
	if ok {
		s.list.Remove(s.list.Size() - 1)
	}
	return
}

func (s *Stack[T]) Peek() (v T, ok bool) {
	return s.list.Get(s.list.Size() - 1)
}
