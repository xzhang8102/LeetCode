package golang

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

// @lc code=start
func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n)
	ans := 0
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			preLen := dp[i-1]
			preMatch := i - preLen - 1
			if preMatch >= 0 && s[preMatch] == '(' {
				dp[i] = dp[i-1] + 2

				if preMatch-1 >= 0 {
					dp[i] += dp[preMatch-1]
				}
			}
			ans = lc32Max(ans, dp[i])
		}
	}
	return ans
}

func lc32Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
