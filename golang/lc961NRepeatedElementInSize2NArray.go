package golang

/*
 * @lc app=leetcode.cn id=961 lang=golang
 *
 * [961] N-Repeated Element in Size 2N Array
 */

// @lc code=start
func repeatedNTimes(nums []int) int {
	for i, num := range nums {
		if i-1 >= 0 && num == nums[i-1] {
			return num
		}
		if i-2 >= 0 && num == nums[i-2] {
			return num
		}
		if i-3 >= 0 && num == nums[i-3] {
			return num
		}
	}
	return -1
}

// @lc code=end
