package golang

/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] Wildcard Matching
 */

// @lc code=start
func isMatch(s string, p string) bool {
	col, row := len(s), len(p)
	dp := make([][]bool, row+1)
	for i := range dp {
		dp[i] = make([]bool, col+1)
	}
	dp[0][0] = true
	for r := 1; r <= row; r++ {
		if p[r-1] == '*' {
			matched := false
			for c := 0; c <= col; c++ {
				if matched || dp[r-1][c] {
					matched = true
					dp[r][c] = true
				}
			}
			continue
		}
		for c := 1; c <= col; c++ {
			if p[r-1] == '?' || s[c-1] == p[r-1] {
				dp[r][c] = dp[r-1][c-1]
			}
		}
	}
	return dp[row][col]
}

// @lc code=end
