package golang

/*
 * @lc app=leetcode.cn id=744 lang=golang
 *
 * [744] Find Smallest Letter Greater Than Target
 */

// @lc code=start
func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	lo, hi := 0, n
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if mid < n && letters[mid] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if hi == n {
		return letters[0]
	} else {
		return letters[hi]
	}
}

// @lc code=end
