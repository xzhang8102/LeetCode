package golang

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	dp := make([][26]int, m+1)
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == byte(j+'a') {
				dp[i][j] = i
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	i, next := 0, 0
	for i < n && next < m {
		next = dp[next][s[i]-'a']
		if s[i] != t[next] {
			return false
		}
		next++
		i++
	}
	return i == n
}

// @lc code=end
