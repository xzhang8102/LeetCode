package golang

import "math"

/*
 * @lc app=leetcode.cn id=821 lang=golang
 *
 * [821] Shortest Distance to a Character
 */

// @lc code=start
func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)
	index := -1
	for i := 0; i < n; i++ {
		switch s[i] {
		case c:
			index = i
		default:
			if index == -1 {
				ans[i] = math.MaxInt64
			} else {
				ans[i] = i - index
			}
		}
	}
	index = n
	for i := n - 1; i >= 0; i-- {
		switch s[i] {
		case c:
			index = i
		default:
			if index < n {
				ans[i] = lc821Min(index-i, ans[i])
			}
		}
	}
	return ans
}

func lc821Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
