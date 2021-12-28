package golang

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	var dp []int
	for _, num := range nums {
		if len(dp) == 0 || num > dp[len(dp)-1] {
			dp = append(dp, num)
		} else {
			for i := 0; i < len(dp); i++ {
				if dp[i] >= num {
					dp[i] = num
					break
				}
			}
		}
	}
	return len(dp)
}

// @lc code=end
