package inwardtraversal

/**
 * Problem:
 * You are given an array of numbers, each representing the height of
 * a vertical line on a graph.
 *
 * A container can be fo~rmed with any pair of these lines,
 * along with the x-axis of the graph.
 *
 * Return the amount of water which the largest container can hold.
 */

/**
 * Solution -> O(n)
 * - Pointers start at opposite ends
 * - Use pointers to calculate the area at the current pointers:
 *		- Height = Min height between the two
 *		- Width = Right index - Left index
 * - If current area is larger than largest, update largest
 * - If minimum height is:
 *		- Left pointer -> Advance left
 *		- Right pointer -> Advance right
 */

func LargestContainer(heights []int) int {
	left, right := 0, len(heights)-1 // index
	largest := 0
	for left < right {
		min_height := heights[left]
		if heights[right] < min_height {
			min_height = heights[right]
		}
		current := min_height * (right - left)
		if current > largest {
			largest = current
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return largest
}
