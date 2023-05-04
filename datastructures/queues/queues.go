package queues

import "github.com/mhrdini/godsa/datastructures/containers"

type Queue[T any] interface {
	containers.Container[T]
	Enqueue(value T)
	Dequeue() (value T, ok bool)
	Peek() (value T, ok bool)
}
