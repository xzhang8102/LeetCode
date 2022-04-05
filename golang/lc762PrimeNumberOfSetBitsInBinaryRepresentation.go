package golang

import "math/bits"

/*
 * @lc app=leetcode.cn id=762 lang=golang
 *
 * [762] Prime Number of Set Bits in Binary Representation
 */

// @lc code=start
func countPrimeSetBits(left int, right int) int {
	ans := 0
	for i := left; i <= right; i++ {
		n := bits.OnesCount(uint(i))
		if lc762IsPrime(n) {
			ans++
		}
	}
	return ans
}

func lc762IsPrime(a int) bool {
	switch a {
	case 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31:
		return true
	default:
		return false
	}
}

// @lc code=end
