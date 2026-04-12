package partiallysortedarrays

/**
 * Problem:
 * A rotated sorted array is an array sorted in ascending order,
 * in which a portion of the array is moved from the beginning to the end.
 * 
 * For example, a possible rotation of [1, 2, 3, 4, 5] is [3, 4, 5, 1, 2],
 * where the first two numbers are moved to the end.
 *
 * Given a rotated sorted array of unique numbers, return the index of a
 * target value. If the target value is not present, return -1.
 */

func FindTargetInRotatedSortedArray(vs []int, target int) int {
	left, right := 0, len(vs) - 1
	for left < right {
		mid := (left + right)/ 2
		if vs[mid] == target {
			return mid
		} else if vs[left] <= vs[mid] {
			if vs[left] <= target && target > vs[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if vs[mid] < target && target <= vs[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if vs != nil && vs[left] == target {
		return left
	}
	return -1
}
