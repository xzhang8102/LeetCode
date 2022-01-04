package golang

/*
 * @lc app=leetcode.cn id=390 lang=golang
 *
 * [390] Elimination Game
 */

// @lc code=start
func lastRemaining(n int) int {
	gap, ans, dir := 1, 1, 1
	for ; n > 1; n /= 2 {
		if dir > 0 {
			ans += gap
			gap *= 2
			dir = -1
		} else {
			if n%2 == 1 {
				ans += gap
			}
			gap *= 2
			dir = 1
		}
	}
	return ans
}

// @lc code=end
