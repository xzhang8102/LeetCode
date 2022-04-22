package golang

/*
 * @lc app=leetcode.cn id=161 lang=golang
 *
 * [161] One Edit Distance
 */

// @lc code=start
func isOneEditDistance(s string, t string) bool {
	n, m := len(s), len(t)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		if i == 0 {
			for j := range dp[i] {
				dp[i][j] = j
			}
		} else {
			dp[i][0] = i
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = lc161Min(dp[i-1][j-1]+1, lc161Min(dp[i-1][j]+1, dp[i][j-1]+1))
			}
		}
	}
	return dp[n][m] == 1
}

func lc161Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
