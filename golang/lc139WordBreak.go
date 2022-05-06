package golang

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func lc139wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, target := range wordDict {
			if i >= len(target) &&
				s[i-len(target):i] == target &&
				dp[i-len(target)] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// @lc code=end
