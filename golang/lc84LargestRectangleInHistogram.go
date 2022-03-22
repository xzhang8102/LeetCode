package golang

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	for _, idx := range stack {
		right[idx] = n
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = lc84Max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func lc84Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
