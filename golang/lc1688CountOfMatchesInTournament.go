package golang

/*
 * @lc app=leetcode.cn id=1688 lang=golang
 *
 * [1688] Count of Matches in Tournament
 */

// @lc code=start
func numberOfMatches(n int) int {
	ans := 0
	for n != 1 {
		ans += n >> 1
		n = (n + 1) / 2
	}
	return ans
}

// @lc code=end
