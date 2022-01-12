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
	leftMin, rightMax := make([]int, n), make([]int, n)
	leftMin[0] = math.MaxInt32
	for i := 1; i < n; i++ {
		leftMin[i] = lc334Min(nums[i-1], leftMin[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = lc334Max(nums[i+1], rightMax[i+1])
	}
	for i := 1; i < n-1; i++ {
		if nums[i] > leftMin[i] && nums[i] < rightMax[i] {
			return true
		}
	}
	return false
}

func lc334Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lc334Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
