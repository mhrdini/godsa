package inwardtraversal

import (
	"sort"
)

/**
 * Problem:
 * Given an array of integers, return all triplets (a, b, c] such that
 * a + b + c = 0. The solution must not contain duplicate triplets
 * (e.g., [1, 2, 3] and [2, 3, 1] are considered duplicate triplets),
 * If no such triptets are found, return an empty array.
 *
 * Each triplet can be arranged in any order, and the output can be returned
 * in any order.
 */

/**
 * Solution: O(n log n + n^2) = O(n^2)
 * - For finding the triplet
 * 		- Sort array -> O(n log n)
 *		- Find sum -> O(n^2)
 * 		- For each of n elements:
 *			- Fix one of the numbers
 *    	- Find the pair of values among the remaining values to its right
 *      	whose sum is equal to the negative of the number: b + c = -a -> O(n)
 * - For avoiding duplicates
 * 		- If pair is not found for the fixed number, advance it until it is a different
 *				number in the list
 */

func TripleSum(arr []int) [][]int {
	sort.Ints(arr)
	n := len(arr)

	var result [][]int

	for i := 0; i < n-2; i++ {
		if arr[i] > 0 {
			break
		}
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		pair := PairSumSorted(arr[i+1:], -arr[i])
		if len(pair) == 2 {
			index_offset := i + 1
			triplet := []int{arr[i], arr[index_offset+pair[0]], arr[index_offset+pair[1]]}
			result = append(result, triplet)
		}
	}

	return result
}
