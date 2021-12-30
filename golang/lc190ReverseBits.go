package golang

/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var ans uint32 = 0
	for i := 0; i < 32 && num > 0; i++ {
		bit := num & 1 << (31 - i) // rightmost bit push to leftmost
		ans |= bit
		num >>= 1
	}
	return ans
}

// @lc code=end
