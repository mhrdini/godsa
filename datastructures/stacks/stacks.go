package stacks

import "github.com/mhrdini/godsa/datastructures/containers"

type Stack[T any] interface {
	containers.Container[T]

	Push(value T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)
}
