package lists

import (
	"github.com/mhrdini/godsa/datastructures/containers"
	"github.com/mhrdini/godsa/datastructures/utils"
)

type List[T any] interface {
	containers.Container[T]
	Reset()
	Sort(comp utils.Comparator[T])
	Add(vs ...T) bool
	InsertAt(i int, vs ...T) bool
	Remove(i int) (T, bool)
	Get(i int) (T, bool)
	Set(i int, v T) bool
	// TODO: Concat(vs ...*List[T])
}
