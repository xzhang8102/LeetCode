package golang

/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] Counting Bits
 */

// @lc code=start
func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ones := 0
		for n := i; n > 0; n = n & (n - 1) {
			ones++
		}
		ans[i] = ones
	}
	return ans
}

// @lc code=end
