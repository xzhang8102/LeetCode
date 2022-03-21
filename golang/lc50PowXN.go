package golang

import "math"

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if x == 0.0 {
		if n < 0 {
			return math.Inf(1)
		}
		return 0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var pow func(x float64, n int) float64
	pow = func(x float64, n int) float64 {
		if n == 1 {
			return x
		}
		m := pow(x, n/2)
		if n&1 == 1 {
			return m * m * x
		}
		return m * m
	}
	return pow(x, n)
}

// @lc code=end
