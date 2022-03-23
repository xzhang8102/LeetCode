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
		start := struct{ r, c int }{up, left}
		for ; start.c < right; start.c++ {
			ans = append(ans, matrix[up][start.c])
		}
		start.c--
		start.r++
		if start.r >= down {
			break
		}
		for ; start.r < down; start.r++ {
			ans = append(ans, matrix[start.r][right-1])
		}
		start.r--
		start.c--
		if start.c < left {
			break
		}
		for ; start.c >= left; start.c-- {
			ans = append(ans, matrix[down-1][start.c])
		}
		start.c++
		start.r--
		if start.r <= up {
			break
		}
		for ; start.r > up; start.r-- {
			ans = append(ans, matrix[start.r][left])
		}
		left++
		right--
		up++
		down--
	}
	return ans
}

// @lc code=end
