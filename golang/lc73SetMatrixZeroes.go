package golang

/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])
	col0 := false
	for _, row := range matrix {
		if row[0] == 0 {
			col0 = true
		}
		for i := 1; i < n; i++ {
			if row[i] == 0 {
				matrix[0][i] = 0
				row[0] = 0
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}

// @lc code=end
