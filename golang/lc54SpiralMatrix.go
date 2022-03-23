package golang

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return nil
	}
	col := len(matrix[0])
	ans := make([]int, 0, row*col)
	for left, right, up, down := 0, col, 0, row; left < right && up < down; {
		for i := left; i < right; i++ {
			ans = append(ans, matrix[up][i])
		}
		for i := up + 1; i < down; i++ {
			ans = append(ans, matrix[i][right-1])
		}
		// make sure there is enough space for further traversal with no repeat
		if right-left > 1 && down-up > 1 {
			for i := right - 2; i >= left; i-- {
				ans = append(ans, matrix[down-1][i])
			}
			for i := down - 2; i > up; i-- {
				ans = append(ans, matrix[i][left])
			}
		}
		left++
		right--
		up++
		down--
	}
	return ans
}

// @lc code=end
