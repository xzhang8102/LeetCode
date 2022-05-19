package golang

import "sort"

/*
 * @lc app=leetcode.cn id=462 lang=golang
 *
 * [462] Minimum Moves to Equal Array Elements II
 */

// @lc code=start
func minMoves2(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
		ans += nums[right] - nums[left]
	}
	return ans
}

// @lc code=end
