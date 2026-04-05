package hashsets

/**
 * Problem:
 * Given a partially completed 9x9 Sudoku board, determine if
 * the current state of the board adheres to the rules of the game:
 *	- Each row and column must contain unique numbers between 1 and 9,
 *		or be empty (represented as 0).
 *	- Each of the nine 3x3 subgrids that compose the grid must
 *		contain unique numbers between 1 and 9, or be empty.
 *  Note: You are asked to determine whether the current state of the board is
 *	valid given these rules, not whether the board is solvable.
 *
 * Constraints:
 *	- Assume each integer on the board falls in the range of (0, 9].
 */

/**
 * Solution -> O(1) since grid size is 9x9, but arbitrary board size O(n^2)
 *	- create a set for every row, column, and sub-grid
 *	- populate as you encounter values, but check first if it exists in the
 *		appropriate row, column, subgrid set
 *	- if it does, then it is a repeated number and therefore is an invalid board
 */

var sudoku_grid_len int = 9
var sudoku_subgrid_len int = 3

func VerifySudokuBoard(board [][]int) bool {
	row_set := make([]map[int]struct{}, sudoku_grid_len)
	col_set := make([]map[int]struct{}, sudoku_grid_len)
	subgrid_set := make([][]map[int]struct{}, sudoku_subgrid_len)

	for i := 0; i < sudoku_grid_len; i++ {
		row_set[i] = make(map[int]struct{})
		col_set[i] = make(map[int]struct{})
	}

	for i := 0; i < sudoku_subgrid_len; i++ {
		subgrid_set[i] = make([]map[int]struct{}, sudoku_subgrid_len)
		for j := 0; j < sudoku_subgrid_len; j++ {
			subgrid_set[i][j] = make(map[int]struct{})
		}
	}

	for i, row := range board {
		for j, col := range row {
			if col == 0 {
				continue
			}

			if _, ok := row_set[i][col]; ok {
				return false
			}

			if _, ok := col_set[j][col]; ok {
				return false
			}

			if _, ok := subgrid_set[i/3][j/3][col]; ok {
				return false
			}

			row_set[i][col] = struct{}{}
			col_set[j][col] = struct{}{}
			subgrid_set[i/3][j/3][col] = struct{}{}
		}
	}

	return true
}

func ExampleSudokuBoard() [][]int {
	return [][]int{
		{3, 0, 6, 0, 5, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{1, 0, 2, 5, 0, 0, 3, 2, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{0, 3, 0, 0, 0, 8, 2, 5, 0},
		{0, 1, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 0, 0, 0},
	}
}
