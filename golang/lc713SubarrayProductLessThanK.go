package golang

/*
 * @lc app=leetcode.cn id=713 lang=golang
 *
 * [713] Subarray Product Less Than K
 */

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {
	ans := 0
	if k <= 1 {
		return ans
	}
	product := 1
	left := 0
	for right, n := range nums {
		product *= n
		for product >= k {
			product /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}

// @lc code=end
