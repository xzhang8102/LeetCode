package golang

/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func missingNumber(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	for _, v := range nums {
		total -= v
	}
	return total
}

// @lc code=end
