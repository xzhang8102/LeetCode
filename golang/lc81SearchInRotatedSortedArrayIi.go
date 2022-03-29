package golang

/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 */

// @lc code=start
func search(nums []int, target int) bool {
	n := len(nums)
	if target == nums[0] || target == nums[n-1] {
		return true
	}
	for lo, hi := 0, n-1; lo <= hi; {
		mid := lo + (hi-lo)>>1
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[lo] && nums[mid] == nums[hi] {
			lo++
			hi--
		} else if nums[lo] <= nums[mid] {
			if target < nums[mid] && nums[lo] <= target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if target > nums[mid] && target < nums[lo] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return false
}

// @lc code=end
