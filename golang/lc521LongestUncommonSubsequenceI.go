package golang

/*
 * @lc app=leetcode.cn id=521 lang=golang
 *
 * [521] Longest Uncommon Subsequence I
 */

// @lc code=start
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return lc521Max(len(a), len(b))
}

func lc521Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
