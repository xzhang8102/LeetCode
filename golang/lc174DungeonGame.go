package golang

/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] Dungeon Game
 */

// @lc code=start
func calculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon), len(dungeon[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			for j := m - 1; j >= 0; j-- {
				if dungeon[i][j] <= 0 {
					if j < m-1 {
						dp[i][j] = dp[i][j+1] - dungeon[i][j]
					} else {
						dp[i][j] = 1 - dungeon[i][j]
					}
				} else {
					if j < m-1 {
						dp[i][j] = dp[i][j+1] - dungeon[i][j]
						if dp[i][j] < 1 {
							dp[i][j] = 1
						}
					} else {
						dp[i][j] = 1
					}
				}
			}
		} else {
			if dungeon[i][m-1] <= 0 {
				dp[i][m-1] = dp[i+1][m-1] - dungeon[i][m-1]
			} else {
				dp[i][m-1] = dp[i+1][m-1] - dungeon[i][m-1]
				if dp[i][m-1] < 1 {
					dp[i][m-1] = 1
				}
			}
			for j := m - 2; j >= 0; j-- {
				if dungeon[i][j] <= 0 {
					dp[i][j] = lc174Min(dp[i+1][j]-dungeon[i][j], dp[i][j+1]-dungeon[i][j])
				} else {
					v1 := dp[i+1][j] - dungeon[i][j]
					if v1 < 1 {
						v1 = 1
					}
					v2 := dp[i][j+1] - dungeon[i][j]
					if v2 < 1 {
						v2 = 1
					}
					dp[i][j] = lc174Min(v1, v2)
				}
			}
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
