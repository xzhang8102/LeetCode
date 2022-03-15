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
	for i, char := range s {
		if i > 0 && char == ')' {
			if s[i-1] == '(' {
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				if i-dp[i-1] >= 1 && s[i-dp[i-1]-1] == '(' {
					if i-dp[i-1] >= 2 {
						dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
			if dp[i] > ans {
				ans = dp[i]
			}
		}
	}
	return ans
}

// @lc code=end
