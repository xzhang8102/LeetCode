package golang

/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] Power of Four
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}

// @lc code=end
