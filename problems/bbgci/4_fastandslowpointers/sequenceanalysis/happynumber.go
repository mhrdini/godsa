package sequenceanalysis

/**
 * Problem:
 * In number theory, .a happy number is defined as a number that,
 * when repeatedly subjected to the process of squaring its digits
 * and summing those squares, eventually leads to 1 [1).
 *
 * An unhappy number will never reach 1 during this process,
 * and will get stuck in an infinite loop.
 *
 * Given an integer, determine if it's a happy number.
 *
 * Example:
 * 23 -> happy
 * 116 -> unhappy
 */

/**
 * Solution:
 *	- Same as a linked list cycle detection problem
 *
 * Time: O(n) where n = length of cycle
 * Space: O(1)
 */

func HappyNumber(v int) bool {

	fast, slow := v, v

	for {
		slow = GetNextNumber(slow)
		fast = GetNextNumber(GetNextNumber(fast))
		if fast == 1 {
			return true
		}
		if fast == slow {
			return false
		}
	}
}

func GetNextNumber(v int) int {
	next := Square(v % 10)
	v %= 10
	for v := v / 10; v != 0; v /= 10 {
		next += Square(v % 10)
	}
	return next
}

func Square(v int) int {
	return v * v
}
