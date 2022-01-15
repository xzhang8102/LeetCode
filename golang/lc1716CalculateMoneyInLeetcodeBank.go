package golang

/*
 * @lc app=leetcode.cn id=1716 lang=golang
 *
 * [1716] Calculate Money in Leetcode Bank
 */

// @lc code=start
func totalMoney(n int) int {
	round, remain := n/7, n%7
	ans := 0
	for i := 1; i <= round; i++ {
		ans += (i + i + 6) * 7 / 2
	}
	for i := round + 1; i <= round+remain; i++ {
		ans += i
	}
	return ans
}

// @lc code=end
