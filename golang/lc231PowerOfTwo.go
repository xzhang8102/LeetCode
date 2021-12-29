package golang

/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] Power of Two
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	mask := 1
	for n != 1 {
		lastBit := n & mask
		if lastBit == 1 {
			return false
		}
		n >>= 1
	}
	return true
}

// @lc code=end
