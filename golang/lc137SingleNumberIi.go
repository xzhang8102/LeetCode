package golang

import "sort"

/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start
func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 3 {
		if i < len(nums)-1 && nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

// @lc code=end
