package golang

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
func findMin(nums []int) int {
	n := len(nums)
	lo, hi := 0, n-1
	rotated := nums[0] > nums[n-1]
	for lo <= hi {
		mid := (lo + hi) >> 1
		prev := mid - 1
		if mid == 0 {
			prev = n - 1
		}
		if nums[prev] >= nums[mid] {
			return nums[mid]
		} else {
			if rotated {
				if nums[mid] >= nums[0] {
					lo = mid + 1
				} else {
					hi = mid - 1
				}
			} else {
				if nums[mid] >= nums[0] {
					hi = mid - 1
				} else {
					lo = mid + 1
				}
			}
		}
	}
	return 0
}

// @lc code=end
