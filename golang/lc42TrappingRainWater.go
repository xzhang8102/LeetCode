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
	left, right := 0, n-1
	leftMax, rightMax := height[left], height[right]
	for left != right {
		if leftMax > rightMax {
			ans += rightMax - height[right]
			right--
			rightMax = lc42Max(rightMax, height[right])
		} else {
			ans += leftMax - height[left]
			left++
			leftMax = lc42Max(leftMax, height[left])
		}
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
