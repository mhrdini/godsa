package hashsets

/**
 * Problem:
 * Find the longest chain of consecutive numbers in an array.
 * Two numbers are consecutive if they have a difference of 1.
 */

/**
 * Solution:
 *  - populate hash set with existing values
 *  - use set to check if a current number is the smallest in the chain:
 *		- i.e. for a current number v, no v-1 exists in the list
 *	- keep incrementing to get the chain length starting from that number,
 *		tracking the longest chain
 */

func LongestChainOfConsecutiveNumbers(vs []int) int {
	set := make(map[int]struct{})
	longest_chain := 0
	for _, v := range vs {
		set[v] = struct{}{}
	}

	for _, v := range vs {
		if _, ok := set[v-1]; ok {
			curr := v
			curr_chain := 1
			for _, ok := set[curr]; ok; {
				curr_chain++
				curr++
				_, ok = set[curr]
			}
			longest_chain = max(longest_chain, curr_chain)
		}
	}
	return longest_chain
}
