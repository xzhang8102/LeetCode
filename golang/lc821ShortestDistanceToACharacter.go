package golang

/*
 * @lc app=leetcode.cn id=821 lang=golang
 *
 * [821] Shortest Distance to a Character
 */

// @lc code=start
func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)
	index := -n
	for i := 0; i < n; i++ {
		if s[i] == c {
			index = i
		}
		ans[i] = i - index
	}
	index = 2 * n
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			index = i
		}
		ans[i] = lc821Min(ans[i], index-i)
	}
	return ans
}

func lc821Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
