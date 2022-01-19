package golang

/*
 * @lc app=leetcode.cn id=1060 lang=golang
 *
 * [1060] Missing Element in Sorted Array
 */

// @lc code=start
func missingElement(nums []int, k int) int {
	n := len(nums)
	if lc1060Missing(n-1, nums) < k {
		return nums[n-1] + k - lc1060Missing(n-1, nums)
	}
	lo, hi := 0, n-1
	for lo != hi {
		mid := lo + (hi-lo)/2
		if lc1060Missing(mid, nums) < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return nums[lo-1] + k - lc1060Missing(lo-1, nums)
}

func lc1060Missing(index int, nums []int) int {
	return nums[index] - nums[0] - index
}

// @lc code=end
