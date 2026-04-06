package inwardtraversal

import (
	"sort"
)

/**
 * Problem:
 * Given an array of integers, return all triplets (a, b, c] such that
 * a + b + c = 0. The solution must not contain duplicate triplets
 * (e.g., [1, 2, 3] and [2, 3, 1] are considered duplicate triplets),
 * If no such triplets are found, return an empty array.
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

func TripleSum(vs []int) [][]int {
	sort.Ints(vs)
	n := len(vs)

	var result [][]int

	for i := 0; i < n-2; i++ {
		if vs[i] > 0 {
			break
		}
		if i > 0 && vs[i] == vs[i-1] {
			continue
		}
		pair := PairSumSorted(vs[i+1:], -vs[i])
		if len(pair) == 2 {
			index_offset := i + 1
			triplet := []int{vs[i], vs[index_offset+pair[0]], vs[index_offset+pair[1]]}
			result = append(result, triplet)
		}
	}

	return result
}
