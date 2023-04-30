package utils

import "golang.org/x/exp/constraints"

// Comparator function needs to return an integer representation of a comparison between x and y:
// - if x < y: return -1
// - if x == y: return 0
// - if x > y: return 1
type Comparator[C any] func(x, y C) int

func OrderedComparator[C constraints.Ordered](x, y C) int {
	if x < y {
		return -1
	} else if x == y {
		return 0
	} else {
		return 1
	}
}
