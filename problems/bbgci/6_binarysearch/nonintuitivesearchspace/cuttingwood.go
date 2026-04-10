package nonintuitivesearchspace

import "github.com/mhrdini/godsa/problems/bbgci/helpers"

/**
 * Problem:
 * You are given an array representing the heights of trees, and an integer k
 * representing the total length of wood that needs to be cut.
 *
 * For this task, a woodcutting machine is set to a certain height, H.
 * The machine cuts off the top part of all trees taller than H,
 * while trees shorter than H remain untouched.
 *
 * Determine the highest possible setting of the woodcutter (H)
 * so that it cuts at least k meters of wood.
 *
 * Assume the woodcutter cannot be set higher than the height of
 * the tallest tree in the array.
 *
 * Constraints:
 *	- It's always possible to attain at least k meters of wood.
 *	- There's at least one tree.
 */

/**
 * Solution:
 *	- Highest setting = Upper bound of h -> Upper bound binary search on h
 *		that satisfies condition of collecting at least k amount of wood
 *  - Look for the upper bound of h that satisfies the condition of cutting at
 *		least k meters of wood
 *
 * Time: O(n log m) where n = # of trees, m = max tree height
 *	- Upper bound binary search among height values [0, m]
 *	- For each iteration, get sum of collected wood by going through n trees
 * Space: O(1)
 */

func CuttingWood(heights []int, k int) int {
	left, right := 0, helpers.Max(heights...)

	for left < right {
		mid := ((left + right) / 2) + 1
		collected := 0

		for _, h := range heights {
			if h > mid {
				collected += h - mid
			}
		}

		if collected >= k {
			left = mid
		} else {
			// k must be less than mid itself, so we cap the upper bound by moving
			// right pointer
			right = mid - 1
		}

	}
	return left
}
