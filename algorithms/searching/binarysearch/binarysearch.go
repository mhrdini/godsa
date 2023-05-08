package binarysearch

import "golang.org/x/exp/constraints"

func Search[T constraints.Ordered](s []T, searchValue T) (idx int, ok bool) {
	if len(s) == 0 {
		return
	}

	lowerBound := 0
	upperBound := len(s) - 1

	for lowerBound <= upperBound {
		midpoint := (upperBound + lowerBound) / 2
		midpointValue := s[midpoint]

		if searchValue == midpointValue {
			idx = midpoint
			ok = true
			return
		} else if searchValue < midpointValue {
			upperBound = midpoint - 1
		} else {
			lowerBound = midpoint + 1
		}
	}

	return
}
