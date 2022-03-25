package golang

/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */

// @lc code=start
func trailingZeroes(n int) int {
	if n < 5 {
		return 0
	}
	// 计算从0 ~ n共有多少个 5^n 的倍数
	// target = 0 + (5 ^ n) * cnt
	// 以130举例，
	// 第一次除以5，得到1-130中有多少个5的倍数
	// 第二次除以5，得到1-130中有多少个25的倍数
	return n/5 + trailingZeroes(n/5)
}

// @lc code=end
