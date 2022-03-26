package golang

/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	row := len(matrix)
	if row == 0 {
		return
	}
	col := len(matrix[0])
	coord := map[int]bool{}
	for r := 0; r < row; r++ {
		clear := false
		for c := 0; c < col; c++ {
			if matrix[r][c] == 0 {
				clear = true
				coord[c] = true
			}
		}
		if clear {
			copy(matrix[r], make([]int, col))
		}
	}
	for c := range coord {
		for r := 0; r < row; r++ {
			matrix[r][c] = 0
		}
	}
}

// @lc code=end
