package inwardtraversal

import (
	"unicode"
)

/**
 * Problem:
 * A palindrome is a sequence of characters that reads the same forward
 * and backward. Given a string, determine if it's a palindrome after
 * removing all non-alphanumeric characters.
 *
 * A character is alphanumeric if it's either a letter or a number.
 *
 * Constraints:
 * The string may include a combination of lowercase English letters,
 * numbers, spaces, and punctuations.
 */

/**
 * Solution: O(n)
 * - Inward traversal skipping over non-alphanum characters
 * - Empty string = true, Single character = true
 * - Return false at first sign where left != right
 */

func IsPalindromeValid(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !IsAlnum(rune(s[left])) {
			left++
		}
		for left < right && !IsAlnum(rune(s[right])) {
			right--
		}
		if rune(s[left]) != rune(s[right]) {
			return false
		}
		if left == right {
			break
		}
		left++
		right--
	}

	return true
}

func IsAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
