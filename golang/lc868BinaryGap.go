package golang

/*
 * @lc app=leetcode.cn id=868 lang=golang
 *
 * [868] Binary Gap
 */

// @lc code=start
func binaryGap(n int) int {
	ans := 0
	last := -1
	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 {
				ans = lc868Max(ans, i-last)
			}
			last = i
		}
		n >>= 1
	}
	return ans
}

func lc868Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
