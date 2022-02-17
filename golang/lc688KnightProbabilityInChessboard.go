package golang

/*
 * @lc app=leetcode.cn id=688 lang=golang
 *
 * [688] Knight Probability in Chessboard
 */

// @lc code=start
func knightProbability(n int, k int, row int, column int) float64 {
	directions := [8][2]int{{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}}
	dp := make([][][]float64, k+1)
	for step := range dp {
		dp[step] = make([][]float64, n)
		for r := 0; r < n; r++ {
			dp[step][r] = make([]float64, n)
			for c := 0; c < n; c++ {
				if step == 0 {
					dp[step][r][c] = 1.0
				} else {
					for _, dir := range directions {
						if newR, newC := r+dir[0], c+dir[1]; newR >= 0 && newR < n && newC >= 0 && newC < n {
							dp[step][r][c] += dp[step-1][newR][newC] / 8.0
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]
}

// @lc code=end
