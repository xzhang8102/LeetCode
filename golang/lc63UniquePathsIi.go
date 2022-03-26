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
	dp := make([]int, col)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if obstacleGrid[r][c] == 1 {
				dp[c] = 0
				continue
			}
			if c > 0 {
				dp[c] += dp[c-1]
			}
		}
	}
	return dp[col-1]
}

// @lc code=end
