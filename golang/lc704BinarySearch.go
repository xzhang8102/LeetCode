package golang

/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func lc704Search(nums []int, target int) int {
	low, hi := 0, len(nums)-1
	for low <= hi {
		mid := (low + hi) >> 1
		found := nums[mid]
		if found == target {
			return mid
		}
		if found < target {
			low = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

// @lc code=end
