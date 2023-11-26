package mergesort

import (
	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](s []T, ascending bool) []T {
	list := append([]T{}, s...)
	mergeSort(list, 0, len(list)-1, ascending)
	return list
}

func mergeSort[T constraints.Ordered](s []T, p, r int, ascending bool) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	mergeSort(s, p, q, ascending)
	mergeSort(s, q+1, r, ascending)
	merge(s, p, q, r, ascending)
}

func merge[T constraints.Ordered](s []T, p, q, r int, ascending bool) {
	llen := q - p + 1 // length of s[p:q+1]
	rlen := r - q     // length of s[q+1:r+1]

	L := append([]T{}, s[p:q+1]...)
	R := append([]T{}, s[q+1:r+1]...)

	i := 0 // smallest remaining element in L
	j := 0 // smallest remaining element in R
	k := p // location in s to fill

	for i < llen && j < rlen {
		switch ascending {
		case true:
			if L[i] <= R[j] {
				s[k] = L[i]
				i++
			} else {
				s[k] = R[j]
				j++
			}
		default:
			if L[i] >= R[j] {
				s[k] = L[i]
				i++
			} else {
				s[k] = R[j]
				j++
			}
		}
		k++
	}

	// after having gone through either L or R entirely,
	// copy remainder of the other to the end of s[p:r]
	for i < llen {
		s[k] = L[i]
		i++
		k++
	}
	for j < rlen {
		s[k] = R[j]
		j++
		k++
	}
}
