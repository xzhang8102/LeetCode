package golang

/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] Number of Longest Increasing Subsequence
 */

// @lc code=start
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	ans := 0
	maxLen := 0
	dp := make([]int, n) // max length of LIS end with number at index i
	count := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			ans = count[i]
		} else if dp[i] == maxLen {
			ans += count[i]
		}
	}
	return ans
}

// @lc code=end
