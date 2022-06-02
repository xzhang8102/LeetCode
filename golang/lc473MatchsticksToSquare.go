package golang

/*
 * @lc app=leetcode.cn id=473 lang=golang
 *
 * [473] Matchsticks to Square
 */

// @lc code=start
func makesquare(matchsticks []int) bool {
	n := len(matchsticks)
	if n < 4 {
		return false
	}
	totalLen := 0
	for _, match := range matchsticks {
		totalLen += match
	}
	if totalLen%4 != 0 {
		return false
	}
	edge := totalLen / 4
	dp := make([]int, 1<<n)
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}
	for state := 1; state < len(dp); state++ {
		for i, match := range matchsticks {
			if state>>i&1 == 0 {
				continue
			}
			prev := state ^ (1 << i)
			if dp[prev] >= 0 && dp[prev]+match <= edge {
				dp[state] = (dp[prev] + match) % edge
				break
			}
		}
	}
	return dp[len(dp)-1] == 0
}

// @lc code=end
