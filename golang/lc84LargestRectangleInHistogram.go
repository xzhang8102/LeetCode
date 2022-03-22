package golang

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

// @lc code=start

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	right := make([]int, n)
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
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
