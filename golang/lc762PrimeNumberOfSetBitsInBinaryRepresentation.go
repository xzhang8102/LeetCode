package golang

/*
 * @lc app=leetcode.cn id=762 lang=golang
 *
 * [762] Prime Number of Set Bits in Binary Representation
 */

// @lc code=start
func countPrimeSetBits(left int, right int) int {
	ans := 0
	for i := left; i <= right; i++ {
		n := lc762CountOnes(uint32(i))
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

func lc762CountOnes(a uint32) int {
	a = a&0x55555555 + (a>>1)&0x55555555
	a = a&0x33333333 + (a>>2)&0x33333333
	a = a&0x0f0f0f0f + (a>>4)&0x0f0f0f0f
	a = a&0x00ff00ff + (a>>8)&0x00ff00ff
	a = a&0x0000ffff + (a>>16)&0x0000ffff
	return int(a)
}

// @lc code=end
