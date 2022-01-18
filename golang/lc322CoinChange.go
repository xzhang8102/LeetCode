package golang

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for remain := 1; remain <= amount; remain++ {
		res := math.MaxInt32
		for _, coin := range coins {
			if remain-coin >= 0 && dp[remain-coin] >= 0 {
				res = lc322Min(res, dp[remain-coin]+1)
			}
		}
		if res != math.MaxInt32 {
			dp[remain] = res
		} else {
			dp[remain] = -1
		}
	}
	return dp[amount]
}

func lc322Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
