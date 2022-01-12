package golang

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	total := m + n - 2
	pick := m - 1
	if n < m {
		pick = n - 1
	}
	ans := 1
	for i, n := total, 0; n < pick && i >= 1; i, n = i-1, n+1 {
		ans *= i
	}
	for i := pick; i > 1; i-- {
		ans /= i
	}
	return ans
}

// @lc code=end
