package golang

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	n := len(nums)
	for ptr, i := 0, 0; i < n; i++ {
		if nums[i] != 0 {
			nums[ptr], nums[i] = nums[i], nums[ptr]
			ptr++
		}
	}
}

// @lc code=end
