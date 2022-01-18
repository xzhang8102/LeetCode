package golang

/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i <= n; i++ {
		max := 0
		for j := i - 1; j >= 1; j-- {
			max = lc343Max(max, lc343Max((i-j)*j, dp[i-j]*j))
		}
		dp[i] = max
	}
	return dp[n]
}

func lc343Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
