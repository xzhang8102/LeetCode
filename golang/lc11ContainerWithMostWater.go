package golang

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	n := len(height)
	ans := 0
	for left, right := 0, n-1; left < right; {
		h := lc11Min(height[left], height[right])
		area := h * (right - left)
		if area > ans {
			ans = area
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func lc11Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
