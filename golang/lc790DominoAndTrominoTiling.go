package golang

/*
 * @lc app=leetcode.cn id=790 lang=golang
 *
 * [790] Domino and Tromino Tiling
 */

// @lc code=start
func numTilings(n int) int {
	const MOD = 1e9 + 7
	// init: when n == 1
	// dp[0] curr col is fully covered
	// dp[1] curr col is covered by its first row
	// dp[2] curr col is covered by its second row
	// dp[3] curr col is empty and prev col is fully covered
	dp := [4]int{1, 0, 0, 1}
	for i := 1; i < n; i++ {
		tmp := [4]int{}
		tmp[0] = (dp[0] + dp[1] + dp[2] + dp[3]) % MOD
		tmp[1] = (dp[3] + dp[2]) % MOD
		tmp[2] = (dp[3] + dp[1]) % MOD
		tmp[3] = dp[0]
		dp = tmp
	}
	return dp[0]
}

// @lc code=end
