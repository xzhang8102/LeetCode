package golang

/*
 * @lc app=leetcode.cn id=2028 lang=golang
 *
 * [2028] Find Missing Observations
 */

// @lc code=start
func missingRolls(rolls []int, mean int, n int) []int {
	offset := 0
	for _, roll := range rolls {
		offset += roll - mean
	}
	if offset < 0 && (mean-6)*n > offset || offset > 0 && (mean-1)*n < offset {
		return nil
	}
	ans := make([]int, n)
	for i := range ans {
		if offset == 0 {
			ans[i] = mean
		} else if offset > 0 {
			if offset <= mean-1 {
				ans[i] = mean - offset
				offset = 0
			} else {
				ans[i] = 1
				offset -= mean - 1
			}
		} else {
			if offset >= mean-6 {
				ans[i] = mean - offset
				offset = 0
			} else {
				ans[i] = 6
				offset += 6 - mean
			}
		}
	}
	return ans
}

// @lc code=end
