package golang

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	ans := 0
	maxPos := 0
	curr := 0 // current location
	next := 0 // next jump start
	for next < len(nums)-1 {
		if curr+nums[curr] > maxPos {
			maxPos = curr + nums[curr]
		}
		if curr == next {
			next = maxPos
			ans++
		}
		curr++
	}
	return ans
}

// @lc code=end
