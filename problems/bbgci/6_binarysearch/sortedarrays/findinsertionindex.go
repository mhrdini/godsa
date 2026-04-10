package sortedarrays

import "golang.org/x/exp/constraints"

/**
 * Problem:
 * You are given a sorted array that contains unique values, along with an
 * integer target.
 *	- If the array contains the target value, return its index.
 *	- Otherwise, return the insertion index. This is the index where the
 *		target would be if it were inserted in order, maintaining the sorted
 *		sequence of the array.
 */

/**
 * Solution:
 *  - Keep getting the midpoint between left and right pointers, shifting the
 *		window according the current midpoint.
 *  - When left and right converge to where target is or should be, left should
 *		be where the index of target should be.
 *
 * Time: O(log n)
 * Space: O(1)
 */
func FindInsertionIndex[T constraints.Ordered](vs []T, target T) int {

	left, right := 0, len(vs)
	for left < right {
		mid := (right + left) / 2
		if target <= vs[mid] {
			// exclude everything larger than mid, i.e. cap the max at mid
			right = mid
		} else {
			// exclude everything equal to and smaller than mid, i.e. cap the min at
			// left + 1
			left = mid + 1
		}
	}
	return left
}
