package hashsets

import "slices"

/**
 * Problem:
 * For each zero in an m x n matrix, set its entire row and column to zero
 * in place.
 */

/**
 * Solution 1:
 *	- Use auxiliary hash sets to keep track of rows and cols with 0s
 *	- 1st pass = Populate sets
 *	- 2nd pass = Use the sets to zero the values on rows and cols with 0s
 *
 * Time: O(m * n) where m = # of rows, n = # of cols
 * Space: O(m + n)
 */

/**
 * Solution 2:
 *	- Use the matrix itself, specifically the first row and first
 *		column to keep track the columns and rows with 0s respectively.
 *	- O(m) to check first row for zeros
 *	- O(n) to check first column for zeros
 *  - O(2 * m * n) = O(m * n) to check through matrix for zeros and set values
 *		where either or both of its row and column has zero
 *
 * Time: O(m) + O(n) + O(m * n) = O(m * n) where m = # of rows, n = # of cols
 * Space: O(1)
 */

func ZeroStripingWithAuxiliary(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	zero_rows := make(map[int]struct{})
	zero_cols := make(map[int]struct{})

	for m, row := range matrix {
		for n, v := range row {
			if v == 0 {
				zero_rows[m] = struct{}{}
				zero_cols[n] = struct{}{}
			}
		}
	}

	for m, row := range matrix {
		for n := range row {
			if _, ok := zero_rows[m]; ok {
				matrix[m][n] = 0
			}
			if _, ok := zero_cols[n]; ok {
				matrix[m][n] = 0
			}
		}
	}
}

func ZeroStriping(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	first_row_zero, first_col_zero := false, false

	// check first row
	if slices.Contains(matrix[0], 0) {
		first_row_zero = true
	}

	// check first column
	for r := range matrix {
		if matrix[r][0] == 0 {
			first_col_zero = true
			break
		}
	}

	m, n := len(matrix), len(matrix[0])

	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0
				matrix[r][0] = 0
			}
		}
	}

	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	if first_row_zero {
		for c := 0; c < n; c++ {
			matrix[0][c] = 0
		}
	}

	if first_col_zero {
		for r := 0; r < m; r++ {
			matrix[r][0] = 0
		}
	}
}

func ExampleMatrix() [][]int {
	return [][]int{
		{1, 2, 3, 4, 5},
		{6, 0, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 0},
	}
}
