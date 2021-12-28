package golang

import "math"

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	ans := 0
	lower := math.MinInt32 / 10
	higher := math.MaxInt32 / 10
	for x != 0 {
		if ans < lower || ans > higher {
			return 0
		}
		digit := x % 10
		ans = ans*10 + digit
		x = x / 10
	}
	return ans
}

// @lc code=end
