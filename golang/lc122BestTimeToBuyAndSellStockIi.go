package golang

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
func lc122maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0] = [2]int{0, -prices[0]}
	for i := 1; i < n; i++ {
		dp[i][0] = lc122Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = lc122Max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[n-1][0]
}

func lc122Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
