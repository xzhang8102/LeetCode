package golang

import "math"

/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt32
	for _, v := range nums[1:] {
		if v > second {
			return true
		}
		if v <= first {
			first = v
		} else {
			second = v
		}
	}
	return false
}

// @lc code=end
