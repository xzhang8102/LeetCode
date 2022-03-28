package golang

/*
 * @lc app=leetcode.cn id=693 lang=golang
 *
 * [693] Binary Number with Alternating Bits
 */

// @lc code=start
func hasAlternatingBits(n int) bool {
	lastBit := -1
	for n > 0 {
		bit := n & 1
		if lastBit == bit {
			return false
		}
		lastBit = bit
		n >>= 1
	}
	return true
}

// @lc code=end
