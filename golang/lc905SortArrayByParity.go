package golang

/*
 * @lc app=leetcode.cn id=905 lang=golang
 *
 * [905] Sort Array By Parity
 */

// @lc code=start
func sortArrayByParity(nums []int) []int {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j]&1 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	return nums
}

// @lc code=end
