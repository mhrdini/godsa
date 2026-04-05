package inwardtraversal

/**
 * Problem:
 * Given an array of integers sorted in ascending order and a target value,
 * return the indexes of any pair of numbers in the array that sum to
 * the target. The order of the indexes in the result doesn't matter.
 * If no pair is found, return an empty array.
 */

/**
 * Solution: O(n)
 * - Use inward traversal
 * - If sum is less than target, advance left pointer
 * - If sum is more than target, advance right pointer
 */

func PairSumSorted(arr []int, target int) []int {
	left := 0
	right := len(arr) - 1

	for sum := arr[left] + arr[right]; left < right; sum = arr[left] + arr[right] {
		if sum == target {
			for right > left+1 && arr[left+1] == arr[left] {
				left++
			}
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}
