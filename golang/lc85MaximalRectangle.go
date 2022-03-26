package golang

/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	row := len(matrix)
	if row == 0 {
		return 0
	}
	ans := 0
	col := len(matrix[0])
	height := make([]int, col)
	for r := 0; r < row; r++ {
		stack := []int{}
		area := 0
		for c := 0; c <= col; c++ {
			if c < col {
				if matrix[r][c] == '1' {
					height[c] += 1
				} else {
					height[c] = 0
				}
			}
			for len(stack) > 0 && (c == col || height[stack[len(stack)-1]] >= height[c]) {
				right := c
				left := -1
				if len(stack) > 1 {
					left = stack[len(stack)-2]
				}
				area = lc85Max(area, (right-left-1)*height[stack[len(stack)-1]])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, c)
		}
		ans = lc85Max(ans, area)
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
