package golang

/*
 * @lc app=leetcode.cn id=713 lang=golang
 *
 * [713] Subarray Product Less Than K
 */

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	left, right := 0, 0
	ans := 0
	product := 1
	for right < n {
		if product*nums[right] < k {
			ans += right - left + 1 // valid subarray end with nums[right]
			product *= nums[right]
			right++
		} else {
			if left == right {
				right++
			} else {
				product /= nums[left]
			}
			left++
		}
	}
	return ans
}

// @lc code=end
