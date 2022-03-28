package golang

/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// @lc code=end
