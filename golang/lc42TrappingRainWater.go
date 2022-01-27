package golang

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	ans := 0
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && h >= height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			width := i - left - 1
			ans += (lc42Min(height[left], h) - height[top]) * width
		}
		stack = append(stack, i)
	}
	return ans
}

func lc42Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lc42Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
