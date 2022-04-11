package golang

/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] Count Numbers with Unique Digits
 */

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	prev, curr := 1, 10
	for i := 2; i <= n; i++ {
		prev, curr = curr, curr+(curr-prev)*(11-i)
	}
	return curr
}

// @lc code=end
