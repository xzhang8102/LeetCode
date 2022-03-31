package golang

/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] Interleaving String
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n3 != n1+n2 {
		return false
	}
	dp := make([]bool, n1+1)
	dp[0] = true
	for j := 1; j <= n1; j++ {
		dp[j] = dp[j-1] && s3[j-1] == s1[j-1]
	}
	for i := 1; i <= n2; i++ {
		dp[0] = dp[0] && s3[i-1] == s2[i-1]
		for j := 1; j <= n1; j++ {
			p3 := i + j - 1
			if dp[j] {
				dp[j] = s3[p3] == s2[i-1]
			}
			if dp[j-1] {
				dp[j] = dp[j] || s3[p3] == s1[j-1]
			}
		}
	}
	return dp[n1]
}

// @lc code=end
