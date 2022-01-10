package golang

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	n := len(nums)
	maxPos := 0
	for i, v := range nums {
		if i <= maxPos {
			if i+v > maxPos {
				maxPos = i + v
			}
			if maxPos >= n-1 {
				return true
			}
		}
	}
	return false
}

// @lc code=end
