package golang

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	for lo, hi := 0, row*col-1; lo <= hi; {
		mid := (lo + hi) >> 1
		r, c := mid/col, mid%col
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return false
}

// @lc code=end
