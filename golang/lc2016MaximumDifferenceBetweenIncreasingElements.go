package golang

/*
 * @lc app=leetcode.cn id=2016 lang=golang
 *
 * [2016] Maximum Difference Between Increasing Elements
 */

// @lc code=start
func maximumDifference(nums []int) int {
	ans := -1
	n := len(nums)
	if n < 2 {
		return -1
	}
	for i, preMin := 1, nums[0]; i < n; i++ {
		if nums[i] > preMin {
			if nums[i]-preMin > ans {
				ans = nums[i] - preMin
			}
		} else {
			preMin = nums[i]
		}
	}
	return ans
}

// @lc code=end
