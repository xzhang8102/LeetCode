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
	dp := make([][]bool, n2+1)
	for i := range dp {
		dp[i] = make([]bool, n1+1)
		if i == 0 {
			dp[0][0] = true
			for j := 1; j <= n1; j++ {
				dp[0][j] = s3[j-1] == s1[j-1] && dp[0][j-1]
			}
		} else {
			dp[i][0] = s3[i-1] == s2[i-1] && dp[i-1][0]
		}
	}
	for p2 := 1; p2 <= n2; p2++ {
		for p1 := 1; p1 <= n1; p1++ {
			p3 := p2 + p1 - 1
			if dp[p2][p1-1] {
				dp[p2][p1] = s3[p3] == s1[p1-1]
			}
			if dp[p2-1][p1] {
				dp[p2][p1] = dp[p2][p1] || s3[p3] == s2[p2-1]
			}
		}
	}
	return dp[len(s2)][len(s1)]
}

// @lc code=end
