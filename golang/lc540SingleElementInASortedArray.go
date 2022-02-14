package golang

/*
 * @lc app=leetcode.cn id=540 lang=golang
 *
 * [540] Single Element in a Sorted Array
 */

// @lc code=start
func singleNonDuplicate(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if mid&1 == 0 {
			if nums[mid] == nums[mid+1] {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			if nums[mid] == nums[mid-1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return nums[left]
}

// @lc code=end
