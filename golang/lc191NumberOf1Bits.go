package golang

/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) int {
	ans := 0
	for num > 0 {
		lastBit := num & 1
		if lastBit == 1 {
			ans++
		}
		num >>= 1
	}
	return ans
}

// @lc code=end
