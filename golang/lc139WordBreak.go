package golang

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dict := map[string]bool{}
	for _, str := range wordDict {
		dict[str] = true
	}
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		valid := false
		for j := i - 1; j >= 0; j-- {
			if dp[j] && dict[s[j:i]] {
				valid = true
				break
			}
		}
		dp[i] = valid
	}
	return dp[n]
}

// @lc code=end
