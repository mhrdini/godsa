package hashsets

/**
 * Problem:
 * For each zero in an m x n matrix, set its entire row and column to zero
 * in place.
 */

func ZeroStriping(matrix [][]int) {
	zero_rows := make(map[int]struct{})
	zero_cols := make(map[int]struct{})

	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				zero_rows[i] = struct{}{}
				zero_cols[j] = struct{}{}
			}
		}
	}

	for i, row := range matrix {
		for j := range row {
			if _, ok := zero_rows[i]; ok {
				matrix[i][j] = 0
			}
			if _, ok := zero_cols[j]; ok {
				matrix[i][j] = 0
			}
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
