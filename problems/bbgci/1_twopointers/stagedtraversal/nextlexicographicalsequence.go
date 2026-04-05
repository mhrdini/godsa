package stagedtraversal

/**
 * Problem:
 * Given a string of lowercase English letters, rearrange the characters
 * to form a new string representing the next immediate sequence in
 * lexicographical (alphabetical) order.
 *
 * If the given string is already last in lexicographical order among
 * all possible arrangements, return the arrangement that's first in
 * lexicographical order.
 *
 * Constraints:
 *	- The string contains at least one character.
 */

/**
* Solution -> O(n)
* - goal: to find the next (i.e. smallest) bigger permutation of the sequence
*	- key info 1: the largest permutation of a sequence will always follow a
*		non-increasing order
* - if a suffix is non-increasing, it is already the maximum permutation,
		i.e. no rearranging will make it bigger
* - so we want to find the shortest suffix that we can rearrange to form
		the next bigger permutation
* - get the left & right indexes, getting each by traversing
*		from end to beginning of sequence:
* 	- get left = the index of first value that, when included, makes the
*			suffix NOT non-increasing
*		- get right = the index of smallest value to the right of left that is
*			larger than the value at left
*	- swap the values in left and right
*		- this is done to get the smallest possible increase
*   - i.e. right is still larger than left, but it is the smallest larger value
*			because we know that the suffix is non-increasing
*	- after swapping, the value at left has been increased minimally
* - but the entire suffix to its right is now non-increasing
*	- and therefore, it is the biggest permutation possible for that suffix
* - however, in order to get the smallest increase of the sequence, we actually
*		need this suffix to be the smallest permutation
* - key info 2: if largest = non-increasing -> increasing = smallest ->
* 	reversing the largest permutation will give us the smallest permutation
*	- so, we reverse the suffix to the right of left
*/

func NextLexicographicalSequence(s string) string {
	rs := []rune(s)
	left, right := len(rs)-2, len(rs)-1

	for left >= 0 && rs[left] >= rs[left+1] {
		left--
	}
	if left >= 0 {
		for right > left && rs[right] <= rs[left] {
			right--
		}
		rs[left], rs[right] = rs[right], rs[left]
	}
	left, right = left+1, len(rs)-1
	for left < right {
		rs[left], rs[right] = rs[right], rs[left]
		left++
		right--
	}
	return string(rs)
}
