package golang

/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] Distinct Subsequences
 */

// @lc code=start
func numDistinct(s string, t string) int {
	n, m := len(s), len(t)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < m; i++ {
		tmp := make([]int, n+1)
		for j := 1; j <= n; j++ {
			if s[j-1] == t[i] {
				tmp[j] = dp[j-1] + tmp[j-1]
			} else {
				tmp[j] = tmp[j-1]
			}
		}
		dp = tmp
	}
	return dp[n]
}

// @lc code=end
