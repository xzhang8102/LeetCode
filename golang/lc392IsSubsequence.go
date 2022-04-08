package golang

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	ptr := 0
	for _, ch := range t {
		if byte(ch) == s[ptr] {
			ptr++
			if ptr == len(s) {
				return true
			}
		}
	}
	return false
}

// @lc code=end
