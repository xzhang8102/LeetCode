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
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = lc42Max(leftMax[i-1], height[i])
	}
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = lc42Max(rightMax[i+1], height[i])
	}
	for i, h := range height {
		ans += lc42Min(leftMax[i], rightMax[i]) - h
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
