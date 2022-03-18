package golang

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
func lc10isMatch(s string, p string) bool {
	row, col := len(s), len(p)
	dp := make([][]bool, row+1)
	for i := range dp {
		dp[i] = make([]bool, col+1)
	}
	// 初始条件，当pattern为空，所有非空的string都应该不能匹配成功
	dp[0][0] = true
	for c := 1; c <= col; c++ {
		// 注意这里r从0开始取值
		for r := 0; r <= row; r++ {
			var char byte
			if r > 0 {
				char = s[r-1]
			}
			// 当前匹配pattern的字符为*时
			// 考虑两种情况：
			// 1. 前一个字符匹配当前字符或者是.
			// 2. 前一个字符不匹配且不是.或者r为0
			// 第二种情况只能放弃匹配*字符，匹配与否取决于p[c-2]时的匹配情况
			// 第一种情况则多一种考虑，看前一个字符的匹配状态
			if p[c-1] == '*' {
				if (p[c-2] == char || p[c-2] == '.') && r > 0 {
					dp[r][c] = dp[r][c-2] || dp[r-1][c]
				} else {
					dp[r][c] = dp[r][c-2]
				}
			} else if p[c-1] == char || p[c-1] == '.' {
				if r > 0 {
					dp[r][c] = dp[r-1][c-1]
				} else {
					dp[r][c] = false
				}
			}
		}
	}
	return dp[row][col]
}

// @lc code=end
