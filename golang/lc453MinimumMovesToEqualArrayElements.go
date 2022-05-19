package golang

/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] Minimum Moves to Equal Array Elements
 */

// @lc code=start
func minMoves(nums []int) int {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	ans := 0
	for _, num := range nums {
		ans += num - min
	}
	return ans
}

// @lc code=end
