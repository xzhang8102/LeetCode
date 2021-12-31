package golang

/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
import "math"

func findPeakElement(nums []int) int {
	n := len(nums)
	lo, hi := 0, n-1
	for lo < hi {
		mid := (lo + hi) >> 1
		prev, next := math.MinInt32, math.MinInt32
		if mid > 0 {
			prev = nums[mid-1]
		}
		if mid < n {
			next = nums[mid+1]
		}
		if nums[mid] > prev && nums[mid] > next {
			return mid
		}
		if nums[mid] <= next {
			lo = mid + 1
		} else if nums[mid] <= prev {
			hi = mid - 1
		}
	}
	return lo
}

// @lc code=end
