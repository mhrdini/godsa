package unidirectionaltraversal

/**
 * Problem:
 * Given an array of integers, modify the array in place to move all zeros
 * to the end while maintaining the relative order of non-zero elements.
 */

func ShiftZerosToEnd(nums []int) {
	left := 0
	for right := range nums {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
