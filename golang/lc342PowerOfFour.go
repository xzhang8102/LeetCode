package golang

/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] Power of Four
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n&1 == 1 {
			return false
		}
		n >>= 1
		if n&1 == 1 {
			return false
		}
		n >>= 1
	}
	return true
}

// @lc code=end
