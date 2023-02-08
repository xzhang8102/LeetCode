package golang

import "strconv"

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func lc9isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// @lc code=end
