package golang

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	n := len(nums)
	left, right := 0, 0
	for right < n {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}

// @lc code=end
