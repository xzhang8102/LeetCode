package golang

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int) {
	size := len(matrix)
	if size <= 1 {
		return
	}
	// reverse every row
	for _, r := range matrix {
		for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
	}
	for r := 0; r < size; r++ {
		for c := 0; c < size-r-1; c++ {
			matrix[r][c], matrix[size-c-1][size-r-1] = matrix[size-c-1][size-r-1], matrix[r][c]
		}
	}
}

// @lc code=end
