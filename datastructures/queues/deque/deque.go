package deque

import "github.com/mhrdini/godsa/datastructures/lists/doublylinkedlist"

type Deque[T any] struct {
	list *doublylinkedlist.List[T]
}

func New[T any](vs ...T) *Deque[T] {
	d := &Deque[T]{list: doublylinkedlist.New[T]()}
	for _, v := range vs {
		d.Enqueue(v)
	}
	return d
}

func (d *Deque[T]) Name() string {
	return "Deque"
}

func (d *Deque[T]) Size() int {
	return d.list.Size()
}

func (d *Deque[T]) Empty() bool {
	return d.list.Empty()
}

func (d *Deque[T]) Values() []T {
	return d.list.Values()
}

func (d *Deque[T]) String() string {
	return d.list.String()
}

func (d *Deque[T]) Reset() {
	d.list.Reset()
}

// Push adds a node to the top of a Deque.
func (d *Deque[T]) Push(v T) {
	d.list.Prepend(v)
}

// Insert adds a node to the bottom of a Deque.
func (d *Deque[T]) Insert(v T) {
	d.list.Append(v)
}

// Pop removes a node from the top of a Deque.
// Returns the removed value (if any) and a boolean determining whether a node was removed.
func (d *Deque[T]) Pop() (v T, ok bool) {
	return d.list.RemoveFront()
}

// Remove removes a node from the bottom of a Deque.
// Returns the removed value (if any) and a boolean determining whether a node was removed.
func (d *Deque[T]) Remove() (v T, ok bool) {
	return d.list.RemoveBack()
}

// Enqueue is an alias for Insert, adding a node to bottom of a Deque.
func (d *Deque[T]) Enqueue(v T) {
	d.Insert(v)
}

// Dequeue is an alias for Pop, removing a node from the top of a Deque.
// Returns the removed value (if any) and a boolean determining whether a node was removed.
func (d *Deque[T]) Dequeue() (v T, ok bool) {
	return d.Pop()
}

func (d *Deque[T]) Peek() (v T, ok bool) {
	return d.list.Get(0)
}
