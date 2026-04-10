package sortedarrays

/**
 * Problem:
 * Given an array of integers sorted in non-decreasing order, return the first
 * and last indexes of a target number.
 *
 * If the target is not found, return [-1, -1].
 */

/**
 * Solution:
 *	- Use binary search twice:
 *		- To find lower bound
 *		- To find upper bound
 *	- When trying to find lower bound:
 *		- Move the right pointer towards mid to converge to a lower bound
 *		- Mid is the value at index (left + right) / 2
 *	- When trying to find upper bound:
 *		- Move the left pointer towards mid to converge to an upper bound
 *		- Mid must be biased towards the right, [(left + right) / 2] + 1
 *      so that when left == mid, the next mid calculated will change so that
 *			the search space will shrink and the loop will eventually end
 *
 * Time: O(log n)
 * Space: O(1)
 */

func FirstAndLastOccurrencesOfNumber(vs []int, target int) (int, int) {
	lowerBound := LowerBoundBinarySearch(vs, target)
	upperBound := UpperBoundBinarySearch(vs, target)
	return lowerBound, upperBound
}

func LowerBoundBinarySearch(vs []int, target int) int {
	left, right := 0, len(vs)-1
	for left < right {
		mid := (right + left) / 2
		switch {
		case vs[mid] < target:
			left = mid + 1
		case target < vs[mid]:
			right = mid - 1
		default:
			right = mid
		}
	}
	if vs != nil && vs[left] == target {
		return left
	}
	return -1
}

func UpperBoundBinarySearch(vs []int, target int) int {
	left, right := 0, len(vs)-1
	for left < right {
		// need to +1 to bias midpoint to the right
		mid := ((right + left) / 2) + 1
		switch {
		case vs[mid] < target:
			left = mid + 1
		case target < vs[mid]:
			right = mid - 1
		default:
			// when doing this when there are 2 elements left,
			// in the case when mid == left,
			// the search space doesn't shrink (i.e. left and right stays the same)
			// and the loop never ends as a result
			// so we need to bias midpoint to the right (as seen in the step above)
			// for upper bound binary search
			left = mid
		}
	}
	if vs != nil && vs[right] == target {
		return right
	}
	return -1
}
