package golang

/*
 * @lc app=leetcode.cn id=790 lang=golang
 *
 * [790] Domino and Tromino Tiling
 */

// @lc code=start
func numTilings(n int) int {
	const MOD = 1e9 + 7
	dp := make([][3]int, n+1)
	dp[0][0] = 1
	dp[1][0] = 1
	for i := 2; i <= n; i++ {
		dp[i][0] = (dp[i-1][0] + dp[i-2][0] + dp[i-1][1] + dp[i-1][2]) % MOD
		dp[i][1] = (dp[i-2][0] + dp[i-1][2]) % MOD
		dp[i][2] = (dp[i-2][0] + dp[i-1][1]) % MOD
	}
	return dp[n][0]
}

// @lc code=end
