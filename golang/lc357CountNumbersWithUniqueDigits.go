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
	if n == 1 {
		return 10
	}
	ans, options := 10, 9
	for i := 0; i <= n-2; i++ {
		options *= 9 - i
		ans += options
	}
	return ans
}

// @lc code=end
