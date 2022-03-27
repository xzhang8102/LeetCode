package golang

/*
 * @lc app=leetcode.cn id=2028 lang=golang
 *
 * [2028] Find Missing Observations
 */

// @lc code=start
func missingRolls(rolls []int, mean int, n int) []int {
	missing := mean * (n + len(rolls))
	for _, v := range rolls {
		missing -= v
	}
	if missing < n || missing > 6*n {
		return nil
	}
	ans := make([]int, n)
	quotient, remainder := missing/n, missing%n
	for i := range ans {
		ans[i] = quotient
		if remainder > 0 {
			ans[i]++
			remainder--
		}
	}
	return ans
}

// @lc code=end
