package golang

import "math"

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	ans := 0
	switch {
	case dividend > 0 && divisor > 0:
		if divisor == 1 {
			return dividend
		}
		for dividend >= divisor {
			ans++
			dividend -= divisor
		}
	case dividend < 0 && divisor < 0:
		if divisor == -1 {
			if dividend == math.MinInt32 {
				return math.MaxInt32
			}
			return -dividend
		}
		for dividend <= divisor {
			ans++
			dividend -= divisor
		}
	case dividend > 0 && divisor < 0:
		if divisor == -1 {
			return -dividend
		}
		divisor = -divisor
		for dividend >= divisor {
			ans++
			dividend -= divisor
		}
		ans *= -1
	case dividend < 0 && divisor > 0:
		if divisor == 1 {
			return dividend
		}
		dividend = -dividend
		for dividend >= divisor {
			ans++
			dividend -= divisor
		}
		ans *= -1
	}
	return ans
}

// @lc code=end
