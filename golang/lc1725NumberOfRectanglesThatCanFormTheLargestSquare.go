package golang

/*
 * @lc app=leetcode.cn id=1725 lang=golang
 *
 * [1725] Number Of Rectangles That Can Form The Largest Square
 */

// @lc code=start
func countGoodRectangles(rectangles [][]int) int {
	maxLen := 0
	ans := 0
	for _, rect := range rectangles {
		if sideLen := lc1725Min(rect[0], rect[1]); sideLen > maxLen {
			ans = 1
			maxLen = sideLen
		} else if sideLen == maxLen {
			ans++
		}
	}
	return ans
}

func lc1725Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
