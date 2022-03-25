package golang

/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */

// @lc code=start
func trailingZeroes(n int) int {
	ans := 0
	for i := 5; i <= n; i += 5 {
		for j := i; j%5 == 0; j /= 5 {
			ans++
		}
	}
	return ans
}

// @lc code=end
