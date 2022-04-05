package golang

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	lo, hi := 1, num
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if num == mid*mid {
			return true
		} else if mid*mid > num {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return false
}

// @lc code=end
