package golang

/*
 * @lc app=leetcode.cn id=1332 lang=golang
 *
 * [1332] Remove Palindromic Subsequences
 */

// @lc code=start
func removePalindromeSub(s string) int {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return 2
		}
	}
	return 1
}

// @lc code=end
