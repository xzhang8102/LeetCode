package golang

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func lc198Rob(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return nums[0]
	}
	first := nums[0]
	second := lc198Max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		first, second = second, lc198Max(second, first+nums[i])
	}
	return second
}

func lc198Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
