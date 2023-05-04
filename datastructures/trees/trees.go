package trees

import (
	"github.com/mhrdini/godsa/datastructures/containers"
)

type Tree[T any] interface {
	containers.Container[T]
}
