package golang

/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	dp := make([]int, col)
	for i := range dp {
		if i == 0 {
			dp[i] = grid[0][0]
		} else {
			dp[i] = dp[i-1] + grid[0][i]
		}
	}
	for r := 1; r < row; r++ {
		for c := 0; c < col; c++ {
			if c == 0 {
				dp[c] += grid[r][0]
			} else {
				dp[c] = grid[r][c] + lc64Min(dp[c-1], dp[c])
			}
		}
	}
	return dp[col-1]
}

func lc64Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
