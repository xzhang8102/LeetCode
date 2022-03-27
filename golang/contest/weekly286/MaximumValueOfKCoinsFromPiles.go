package weekly286

func maxValueOfCoins(piles [][]int, k int) int {
	pilesNum := len(piles)
	// dp[i][j] 0到i组，取j次 可以获得的最大价值
	dp := make([][]int, pilesNum)
	for i := range dp {
		dp[i] = make([]int, k+1)
		if i == 0 {
			for j := 1; j <= len(piles[0]) && j <= k; j++ {
				dp[0][j] = dp[0][j-1] + piles[0][j-1]
			}
		}
	}
	for i := 1; i < pilesNum; i++ {
		// 取的总个数 - j
		for j := 1; j <= k; j++ {
			sum := 0
			// 第i组取的个数 - m
			for m := 0; m <= len(piles[i]); m++ {
				if m >= 1 {
					sum += piles[i][m-1]
				}
				if j >= m {
					dp[i][j] = max(dp[i][j], dp[i-1][j-m]+sum)
				}
			}
		}
	}
	return dp[pilesNum-1][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
