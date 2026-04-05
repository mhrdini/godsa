package inwardtraversal

import "fmt"

/**
 * Problem:
 * Given an array of integers sorted in ascending order and a target value,
 * return the indexes of any pair of numbers in the array that sum to
 * the target. The order of the indexes in the result doesn't matter.
 * If no pair is found, return an empty array.
 */

func PairSumSorted(arr []int, target int) []int {
	left := 0
	right := len(arr) - 1

	for sum := arr[left] + arr[right]; sum != target; sum = arr[left] + arr[right] {
		if sum < target {
			left++
		} else {
			right--
		}
	}

	pair := []int{left, right}

	fmt.Println(pair)

	return pair
}
