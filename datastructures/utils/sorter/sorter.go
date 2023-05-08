package sorter

import (
	"sort"

	"github.com/mhrdini/godsa/datastructures/utils/comparator"
)

func Sort[T any](vs []T, comp comparator.Comparator[T]) {
	sort.Sort(sortable[T]{values: vs, comparator: comp})
}

type sortable[T any] struct {
	values     []T
	comparator comparator.Comparator[T]
}

func (s sortable[T]) Len() int {
	return len(s.values)
}

func (s sortable[T]) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}
func (s sortable[T]) Less(i, j int) bool {
	return s.comparator(s.values[i], s.values[j]) < 0
}
