package quicksort

import (
	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](s []T, ascending bool) []T {
	list := append([]T{}, s...)
	quickSort(list, 0, len(list)-1, ascending)
	return list
}

func quickSort[T constraints.Ordered](s []T, p, r int, ascending bool) {
	if p < r {
		q := partition(s, p, r, ascending)
		quickSort(s, p, q-1, ascending)
		quickSort(s, q+1, r, ascending)
	}
}

func partition[T constraints.Ordered](s []T, p, r int, ascending bool) int {
	x := s[r]
	i := p - 1
	for j := p; j < r; j++ {
		switch ascending {
		case true:
			if s[j] <= x {
				i++
				s[i], s[j] = s[j], s[i]
			}
		default:
			if s[j] >= x {
				i++
				s[i], s[j] = s[j], s[i]
			}
		}

	}
	s[i+1], s[r] = s[r], s[i+1]
	return i + 1
}
