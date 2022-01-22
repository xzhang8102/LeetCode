package golang

/*
 * @lc app=leetcode.cn id=1332 lang=golang
 *
 * [1332] Remove Palindromic Subsequences
 */

// @lc code=start
func removePalindromeSub(s string) int {
	for i, n := 0, len(s); i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return 2
		}
	}
	return 1
}

// @lc code=end
