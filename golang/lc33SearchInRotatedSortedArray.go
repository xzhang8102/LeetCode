package golang

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
	n := len(nums)
	ans := -1
	for lo, hi := 0, n-1; lo <= hi; {
		mid := (lo + hi) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[0] { // in first half
			if nums[0] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if nums[n-1] >= target && target > nums[mid] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return ans
}

// @lc code=end
