package golang

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	if n == 0 {
		return false
	}
	for i := n - 1; i >= 0; i-- {
		next := i + nums[i]
		dp[i] = next >= n-1
		if !dp[i] {
			for ; next > i; next-- {
				if dp[next] {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[0]
}

// @lc code=end
