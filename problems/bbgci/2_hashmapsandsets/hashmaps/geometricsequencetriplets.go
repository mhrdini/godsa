package hashmaps

/**
 * Problem:
 * A geometric sequence triplet is a sequence of three numbers
 * where each successive number is obtained by multiplying the
 * preceding number by a constant called the common ratio.
 *
 * Let's examine three triplets to understand how this works:
 *  - (1, 2, 4): Geometric sequence with a ratio of 2
 * 		(i.e., [1, 1·2 = 2, 2·2 = 4]).
 * 	- (5, 15, 45): Geometric sequence with a ratio of 3
 * 		(i.e., [5, 5·3 = 15, 15·3 = 45]).
 * 	- (2, 3, 4): Not a geometric sequence.
 *
 * Given an array of integers and a common ratio r,
 * find all triplets of indexes (i, j, k) that
 * follow a geometric sequence for i < j < k.
 * It's possible to encounter duplicate triplets in the array.
 */

func GeometricSequenceTriplet(vs []int, ratio int) int {
	count := make(map[int]int) // count key = first element
	pairs := make(map[int]int) // pair key = second element
	triplets := 0

	for _, v := range vs {
		if v%ratio == 0 {
			// if it's the third element, then form triplets with the pairs whose
			// second element is itself divided by the ratio
			triplets += pairs[v/ratio]
			// if it's second element, then form pairs with the first element in
			// sequence, which is itself divided by the ratio
			pairs[v] += count[v/ratio]
		}
		count[v]++
	}

	return triplets
}
