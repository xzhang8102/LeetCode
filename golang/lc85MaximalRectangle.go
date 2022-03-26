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
		left, right := make([]int, col), make([]int, col)
		for c := 0; c < col; c++ {
			if matrix[r][c] == '1' {
				height[c] += 1
			} else {
				height[c] = 0
			}
			for len(stack) > 0 && height[stack[len(stack)-1]] >= height[c] {
				right[stack[len(stack)-1]] = c
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				left[c] = -1
			} else {
				left[c] = stack[len(stack)-1]
			}
			stack = append(stack, c)
		}
		for _, idx := range stack {
			right[idx] = col
		}
		area := 0
		for i := 0; i < col; i++ {
			area = lc85Max(area, (right[i]-left[i]-1)*height[i])
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
