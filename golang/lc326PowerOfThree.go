package golang

/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] Power of Three
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n <= 0 {
		return false
	}
	if n%3 == 0 {
		return isPowerOfThree(n / 3)
	} else {
		return false
	}
}

// @lc code=end
