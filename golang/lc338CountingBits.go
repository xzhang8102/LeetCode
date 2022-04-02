package golang

/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] Counting Bits
 */

// @lc code=start
func countBits(n int) []int {
	ans := make([]int, n+1)
	highest := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highest = i
		}
		ans[i] = ans[i-highest] + 1
	}
	return ans
}

// @lc code=end
