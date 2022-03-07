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
	size := 32 << (^uint(0) >> 63)
	negatvie := ((dividend ^ divisor) >> (size - 1) & 0x1) == 1
	dividend = lc29Abs(dividend)
	divisor = lc29Abs(divisor)
	for dividend >= divisor {
		i := 1
		for tmp := divisor; dividend >= tmp; i, tmp = i<<1, tmp<<1 {
			dividend -= tmp
			ans += i
		}
	}
	if negatvie {
		ans *= -1
	}
	if ans > math.MaxInt32 {
		return math.MaxInt32
	}
	if ans < math.MinInt32 {
		return math.MinInt32
	}
	return ans
}

func lc29Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
