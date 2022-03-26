package golang

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	if row == 0 {
		return 0
	}
	col := len(obstacleGrid[0])
	dp := make([][]int, row)
	for i := range obstacleGrid {
		dp[i] = make([]int, col)
		if i == 0 {
			dp[0][0] = 1
			if obstacleGrid[0][0] == 1 {
				dp[0][0] = 0
			}
			for j := 1; j < col; j++ {
				if obstacleGrid[0][j] == 1 {
					dp[0][j] = 0
				} else {
					dp[0][j] = dp[0][j-1]
				}
			}
		} else {
			if obstacleGrid[i][0] == 1 {
				dp[i][0] = 0
			} else {
				dp[i][0] = dp[i-1][0]
			}
		}
	}
	for r := 1; r < row; r++ {
		for c := 1; c < col; c++ {
			if obstacleGrid[r][c] == 1 {
				dp[r][c] = 0
			} else {
				dp[r][c] = dp[r-1][c] + dp[r][c-1]
			}
		}
	}
	return dp[row-1][col-1]
}

// @lc code=end
