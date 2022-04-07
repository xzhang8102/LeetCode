package golang

/*
 * @lc app=leetcode.cn id=796 lang=golang
 *
 * [796] Rotate String
 */

// @lc code=start
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	n := len(s)
	for shift := 0; shift < n; shift++ {
		if goal == s[shift:]+s[:shift] {
			return true
		}
	}
	return false
}

// @lc code=end
