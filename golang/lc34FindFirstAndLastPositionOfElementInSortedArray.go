package golang

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	// find the lower bound
	left := lc34Helper(0, len(nums), target, nums)
	// find the higher bound
	right := lc34Helper(0, len(nums), target+1, nums)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	return []int{left, right - 1}
}

func lc34Helper(lo, hi, target int, nums []int) int {
	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// @lc code=end
