package golang

import "math"

/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] Dungeon Game
 */

// @lc code=start
func calculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon), len(dungeon[0])
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt64
		}
	}
	dp[n][m-1], dp[n-1][m] = 1, 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			dp[i][j] = lc174Max(1, lc174Min(dp[i+1][j], dp[i][j+1])-dungeon[i][j])
		}
	}
	return dp[0][0]
}

func lc174Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lc174Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
