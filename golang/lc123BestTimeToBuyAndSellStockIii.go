package golang

/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 */

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2][3]int, n)
	// dp[i][j][k] - the ith day; current share 0/1; trade counts 0/1/2
	dp[0][0][1] = 0
	dp[0][0][2] = 0
	dp[0][1][0] = -prices[0]
	dp[0][1][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0][1] = lc123Max(dp[i-1][0][1], dp[i-1][1][0]+prices[i])
		dp[i][0][2] = lc123Max(dp[i-1][0][2], dp[i-1][1][1]+prices[i])
		dp[i][1][0] = lc123Max(dp[i-1][1][0], -prices[i])
		dp[i][1][1] = lc123Max(dp[i-1][1][1], dp[i-1][0][1]-prices[i])
	}
	return lc123Max(0, lc123Max(dp[n-1][0][1], dp[n-1][0][2]))
}

func lc123Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
