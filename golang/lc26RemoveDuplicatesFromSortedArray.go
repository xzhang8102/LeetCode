package golang

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func lc26removeDuplicates(nums []int) int {
	n := len(nums)
	slow, fast := 0, 1
	for fast < n {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		} else {
			fast++
		}
	}
	return slow + 1
}

// @lc code=end
