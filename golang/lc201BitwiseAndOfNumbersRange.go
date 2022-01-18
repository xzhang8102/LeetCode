package golang

/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 */

// @lc code=start
func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	for left < right {
		left, right = left>>1, right>>1
		shift++
	}
	return left << shift
}

// @lc code=end
