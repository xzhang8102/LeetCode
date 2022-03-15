package golang

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i] = (digits[i] + 1) % 10
		if digits[i] != 0 {
			return digits
		}
	}
	tmp := make([]int, n+1)
	tmp[0] = 1
	return tmp
}

// @lc code=end
