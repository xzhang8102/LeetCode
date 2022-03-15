package golang

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	ans := -1
	for lo, hi := 0, x; lo <= hi; {
		mid := lo + (hi-lo)>>1
		if mid*mid <= x {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

// @lc code=end
