package golang

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	n := len(s)
	dp := [][]bool{}
	for i := 0; i < n; i++ {
		row := make([]bool, n)
		row[i] = true
		dp = append(dp, row)
	}
	maxLen := 1
	begin := 0
	for currLen := 2; currLen <= n; currLen++ {
		for i := 0; i+currLen-1 < n; i++ {
			valid := s[i] == s[i+currLen-1]
			lo, hi := i+1, i+currLen-2
			if lo <= hi {
				valid = valid && dp[lo][hi]
			}
			dp[i][i+currLen-1] = valid
			if valid && currLen > maxLen {
				maxLen = currLen
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

// @lc code=end
