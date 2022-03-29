package golang

/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] Distinct Subsequences
 */

// @lc code=start
func numDistinct(s string, t string) int {
	n, m := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		if i == 0 {
			for j := range dp[i] {
				dp[i][j] = 1
			}
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[j-1] == t[i-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

// @lc code=end
