package golang

/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	y := len(matrix)
	if y == 0 {
		return 0
	}
	x := len(matrix[0])
	left := make([][]int, y) // 记录每个点向左最多连续的1的个数
	for i, row := range matrix {
		left[i] = make([]int, x)
		for j, v := range row {
			if v == '1' {
				if j == 0 {
					left[i][j] = 1
				} else {
					left[i][j] = left[i][j-1] + 1
				}
			}
		}
	}
	ans := 0
	for r := 0; r < y; r++ {
		for c := 0; c < x; c++ {
			if matrix[r][c] == '1' {
				area := left[r][c]
				width := left[r][c]
				for i := r - 1; i >= 0 && left[i][c] > 0; i-- {
					width = lc85Min(width, left[i][c])
					area = lc85Max(area, (r-i+1)*width)
				}
				ans = lc85Max(ans, area)
			}
		}
	}
	return ans
}

func lc85Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lc85Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
