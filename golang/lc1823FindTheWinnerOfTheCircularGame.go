package golang

/*
 * @lc app=leetcode.cn id=1823 lang=golang
 *
 * [1823] Find the Winner of the Circular Game
 */

// @lc code=start
func findTheWinner(n int, k int) int {
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (ans + k) % i
	}
	return ans + 1
}

// @lc code=end
