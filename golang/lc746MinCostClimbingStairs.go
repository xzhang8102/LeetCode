package golang

/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = lc746Min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func lc746Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
