package golang

/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start
func findDuplicate(nums []int) int {
	fast, slow := nums[nums[0]], nums[0]
	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	slow = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return fast
}

// @lc code=end
