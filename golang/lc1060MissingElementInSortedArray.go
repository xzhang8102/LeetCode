package golang

/*
 * @lc app=leetcode.cn id=1060 lang=golang
 *
 * [1060] Missing Element in Sorted Array
 */

// @lc code=start
func missingElement(nums []int, k int) int {
	n := len(nums)
	total := nums[n-1] - nums[0] + 1 - n
	if total < k {
		return nums[n-1] + k - total
	}
	for lo, hi := 0, n-1; lo <= hi; {
		mid := lo + (hi-lo)/2
		curr := nums[mid] - nums[0] - mid
		if mid == 0 {
			return nums[0] + k
		}
		prev := nums[mid-1] - nums[0] - mid + 1
		if k > prev && k <= curr {
			return nums[mid-1] + k - prev
		} else if k <= prev {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}

// @lc code=end
