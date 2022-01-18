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
	memo := make([]int, amount+1)
	var dp func(int) int
	dp = func(remain int) int {
		if remain < 0 {
			return -1
		}
		if remain == 0 {
			return 0
		}
		if memo[remain] != 0 {
			return memo[remain]
		}
		min := math.MaxInt32
		for _, val := range coins {
			res := dp(remain - val)
			if res >= 0 && res+1 < min {
				min = res + 1
			}
		}
		if min != math.MaxInt32 {
			memo[remain] = min
		} else {
			memo[remain] = -1
		}
		return memo[remain]
	}
	return dp(amount)
}

// @lc code=end
