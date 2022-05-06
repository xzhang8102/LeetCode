package golang

import "strings"

/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] Word Break II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	dict := map[string]bool{}
	for _, word := range wordDict {
		dict[word] = true
	}
	n := len(s)
	dp := make([][][]string, n)
	var backtrack func(start int) [][]string
	backtrack = func(start int) [][]string {
		if dp[start] != nil {
			return dp[start]
		}
		ret := [][]string{}
		for i := start; i < n; i++ {
			word := s[start : i+1]
			if dict[word] {
				if i == n-1 {
					ret = append(ret, []string{word})
				} else {
					for _, words := range backtrack(i + 1) {
						ret = append(ret, append([]string{word}, words...))
					}
				}
			}
		}
		dp[start] = ret
		return ret
	}
	ans := []string{}
	backtrack(0)
	for _, words := range dp[0] {
		ans = append(ans, strings.Join(words, " "))
	}
	return ans
}

// @lc code=end
