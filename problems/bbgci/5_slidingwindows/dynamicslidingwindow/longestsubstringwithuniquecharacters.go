package dynamicslidingwindow

/**
 * Problem:
 * Given a string, determine the length of the longest substring that consists
 * only of unique substrings.
 */

/**
 * Solution:
 *  - Keep dynamic sliding window that:
 *		- Shrinks (i.e. moves the left) to the letter after the last occurrence
 *			of the letter at the right pointer, effectively excluding that last
 *			occurrence
 *		- Updates the last seen index of letter at right pointer
 *		- Moves the right pointer to the right by 1 each time
 *
 * Time: O(n)
 * Space: O(m) where m is total # of unique characters encountered
 */

func LongestSubstringWithUniqueCharacters(s string) int {
	seen := make(map[byte]int)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		// shrink to exclude the last occurrence of the rightmost character
		// if it has occurred before
		if idx, ok := seen[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		seen[s[right]] = right
		maxLen = max(maxLen, right+1-left)
	}

	return maxLen
}
