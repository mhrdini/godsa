package fixedslidingwindow

/**
 * Problem:
 * Given two strings, s and t, both consisting of lowercase English letters,
 * return the number of substrings in s that are anagrams of t.
 *
 * An anagram is a word or phrase formed by rearranging the letters of
 * another word or phrase, using all the original letters exactly once.
 */

/**
 * Solution:
 *	- Populate a 26-length array, one for each letter, of the original t string
 *  - Keep and update another 26-length array for the frequencies of letters of
 *		the current fixed sliding window
 *  - If the current sliding window has the same frequencies as the original
 *		string's letter frequencies, update total count of anagrams
 *
 * Time: O(n)
 * Space: O(1)
 */

func SubstringAnagrams(s, t string) int {
	count := 0

	// anagram = order of letters doesn't matter, frequency does
	expected_freqs := make([]int, 26)

	for _, c := range t {
		idx := getZeroBasedLetterIndex(c) // a = 0, z = 25
		expected_freqs[idx]++
	}

	window_freqs := make([]int, 26)
	left, right := 0, 0

	for right < len(s) {
		window_freqs[getZeroBasedLetterIndex(rune(s[right]))]++

		// anagrams of t will have same length of t, so when the sliding window has
		// reached that length, check if it is an anagram
		// update count if it is, in any case slide the window, and update
		// the window freqs
		if right-left+1 == len(t) {
			if same(window_freqs, expected_freqs) {
				count++
			}
			window_freqs[getZeroBasedLetterIndex(rune(s[left]))]--
			left++
		}
		right++
	}

	return count
}

func getZeroBasedLetterIndex(c rune) int {
	return int(c - 'a')
}

func same(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
