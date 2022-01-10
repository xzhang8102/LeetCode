package golang

/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	dp1, dp2 := make([]int, n), make([]int, n)
	dp1[0] = 0
	for i := 1; i < n; i++ {
		if i == 1 {
			dp1[i] = nums[1]
		} else {
			dp1[i] = lc213Max(nums[i]+dp1[i-2], dp1[i-1])
		}
	}
	dp2[0] = nums[0]
	for i := 1; i < n; i++ {
		if i == 1 || i == n-1 {
			dp2[i] = dp2[i-1]
		} else {
			dp2[i] = lc213Max(nums[i]+dp2[i-2], dp2[i-1])
		}
	}
	return lc213Max(dp2[len(dp2)-1], dp1[len(dp1)-1])
}

func lc213Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
