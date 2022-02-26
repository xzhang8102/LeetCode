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
	for left, right := 0, 0; right < n; {
		if left == right || nums[right] > nums[left] {
			if right > left && nums[right]-nums[left] > ans {
				ans = nums[right] - nums[left]
			}
			right++
		} else {
			left++
		}
	}
	return ans
}

// @lc code=end
