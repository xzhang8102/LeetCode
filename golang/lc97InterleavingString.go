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
	dp := make([][]int, n2+1)
	for i := range dp {
		dp[i] = make([]int, n1+1)
	}
	var dfs func(int, int, int)
	dfs = func(p1, p2, p3 int) {
		if p1 == len(s1) && p2 == len(s2) && p3 == len(s3) {
			dp[p2][p1] = 1
			return
		}
		if dp[p2][p1] != 0 {
			return
		}
		if p1 < len(s1) && s1[p1] == s3[p3] {
			dfs(p1+1, p2, p3+1)
			if dp[p2][p1+1] == 1 {
				dp[p2][p1] = 1
			}
		}
		if p2 < len(s2) && s2[p2] == s3[p3] {
			dfs(p1, p2+1, p3+1)
			if dp[p2+1][p1] == 1 {
				dp[p2][p1] = 1
			}
		}
		if dp[p2][p1] != 1 {
			dp[p2][p1] = -1
		}
	}
	dfs(0, 0, 0)
	return dp[len(s2)][len(s1)] == 1
}

// @lc code=end
