package golang

/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	curr, prev := 0, 0
	for i := 2; i <= n; i++ {
		tmp := curr
		curr = lc746Min(curr+cost[i-1], prev+cost[i-2])
		prev = tmp
	}
	return curr
}

func lc746Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
