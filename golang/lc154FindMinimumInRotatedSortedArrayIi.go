package golang

/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] Find Minimum in Rotated Sorted Array II
 */

// @lc code=start
func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if nums[mid] < nums[hi] {
			hi = mid
		} else if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi--
		}
	}
	return nums[hi]
}

// @lc code=end
