package hashmaps

/**
 * Problem:
 * Given an array of integers, return the indexes of any two numbers that
 * add up to a target. The order of the indexes in the result doesn't matter.
 * If no pair is found, return an empty array.
 *
 * Constraints:
 * - The same index cannot be used twice in the result.
 */

/**
 * Solution 1:
 *  - two passes
 *	- first construct a map of each value to its index
 *	- then for each value:
 *		- calculate its complement, i.e. the other value needed to get the sum
 *		- if the complement exists in the map i.e. in the array, then
 *			return the index of the value and the index of the complement
 *
 * Time: O(n)
 * Space: O(1)
 */

/**
 * Solution 2:
 *  - one pass
 *	- similar, but populate as you go
 *
 * Time: O(n)
 * Space: O(1)
 */

func PairSumUnsortedTwoPass(vs []int, target int) []int {

	complement_idx := map[int]int{}

	for i, v := range vs {
		complement_idx[v] = i
	}

	for i, v := range vs {
		complement := target - v
		idx, ok := complement_idx[complement]
		if ok {
			return []int{i, idx}
		}
	}
	return []int{}
}

func PairSumUnsorted(vs []int, target int) []int {
	complement_idx := map[int]int{}

	for i, v := range vs {
		complement_idx[v] = i
		complement := target - v
		idx, ok := complement_idx[complement]
		if ok && i != idx {
			return []int{idx, i}
		}
	}
	return []int{}
}
