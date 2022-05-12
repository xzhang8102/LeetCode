package golang

/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] Best Time to Buy and Sell Stock IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	if k == 0 {
		return 0
	}
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, k+1)
	for i := range dp {
		dp[i][1] = -prices[0]
	}
	for i := 1; i < n; i++ {
		tmp := make([][2]int, k+1)
		tmp[0][1] = lc188Max(-prices[i], dp[0][1])
		for j := 1; j <= k; j++ {
			tmp[j][0] = lc188Max(dp[j][0], dp[j-1][1]+prices[i])
			tmp[j][1] = lc188Max(dp[j][1], dp[j][0]-prices[i])
		}
		dp = tmp
	}
	ans := 0
	for i := 0; i <= k; i++ {
		ans = lc188Max(ans, dp[i][0])
	}
	return ans
}

func lc188Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
