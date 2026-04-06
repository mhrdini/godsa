package unidirectionaltraversal

/**
 * Problem:
 * Given an array of integers, modify the array in place to move all zeros
 * to the end while maintaining the relative order of non-zero elements.
 */

func ShiftZerosToEnd(vs []int) {
	left := 0
	for right := range vs {
		if vs[right] != 0 {
			vs[left], vs[right] = vs[right], vs[left]
			left++
		}
	}
}
