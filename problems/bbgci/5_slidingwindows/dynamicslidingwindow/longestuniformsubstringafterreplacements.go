package dynamicslidingwindow

import "strings"

/**
 * Problem:
 * A uniform substring is one in which all characters are identical.
 * Given a string, determine the length of the longest uniform substring
 * that can be formed by replacing up to k characters.
 */

/**
 * Solution:
 *	- For the current window, track the frequency of the most frequent character
 *  - The difference b/w the window size and this max frequency will be the
 *    number of characters you need to replace to make it uniform(ly that most
 *		frequent character)
 *  - The condition that the window must maintain before expanding is that this
 *		# of characters to be replace must not be larger than k
 *  - Therefore we keep shrinking (i.e. moving the left pointer by 1) until this
 *    condition is satisfied again, and certainly before we start expanding and
 *		finding longer uniform substrings
 *
 * Time: O(n)
 * Space: O(m) where m is # of unique chars encountered
 */

func LongestUniformSubstringAfterReplacements(s string, k int) (string, int) {
	count := make(map[byte]int)
	left := 0
	maxFreq := 0
	maxLen := 0
	longest := ""

	for right := 0; right < len(s); right++ {
		// update the count since we're expanding
		count[s[right]]++

		// get the most frequent char in window
		maxFreq = max(maxFreq, count[s[right]])

		// if invalid, shrink
		// # of chars to replace = length of window - frequency of most freq char in window
		for (right+1-left)-maxFreq > k {
			// update the count since we're shrinking
			count[s[left]]--
			left++
		}

		maxLen = max(maxLen, right+1-left)
		longest = max(longest, strings.Repeat(string(s[left]), right+1-left))
	}

	return longest, maxLen
}
