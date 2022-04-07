package golang

import "strings"

/*
 * @lc app=leetcode.cn id=796 lang=golang
 *
 * [796] Rotate String
 */

// @lc code=start
func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}

// @lc code=end
