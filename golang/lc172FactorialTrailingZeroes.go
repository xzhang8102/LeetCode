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
	return n/5 + trailingZeroes(n/5)
}

// @lc code=end
