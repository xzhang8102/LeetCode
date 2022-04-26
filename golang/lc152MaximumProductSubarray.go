package golang

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
func maxProduct(nums []int) int {
	n := len(nums)
	// max/min product of the sub-array ends with nums[i]
	dp := make([][2]int, n)
	dp[0] = [2]int{nums[0], nums[0]}
	ans := nums[0]
	for i := 1; i < n; i++ {
		dp[i][0] = lc152Max(nums[i], lc152Max(dp[i-1][0]*nums[i], dp[i-1][1]*nums[i]))
		dp[i][1] = lc152Min(nums[i], lc152Min(dp[i-1][0]*nums[i], dp[i-1][1]*nums[i]))
		ans = lc152Max(ans, dp[i][0])
	}
	return ans
}

func lc152Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lc152Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
