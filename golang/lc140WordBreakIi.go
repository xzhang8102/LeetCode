package golang

import "strings"

/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] Word Break II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	ans := []string{}
	dict := map[string]bool{}
	for _, word := range wordDict {
		dict[word] = true
	}
	n := len(s)
	pick := []int{0}
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == n {
			tmp := make([]string, 0, len(pick))
			for i := 1; i < len(pick); i++ {
				tmp = append(tmp, s[pick[i-1]:pick[i]])
			}
			ans = append(ans, strings.Join(tmp, " "))
			return
		}
		for i := start; i < n; i++ {
			if dict[s[start:i+1]] {
				pick = append(pick, i+1)
				backtrack(i + 1)
				pick = pick[:len(pick)-1]
			}
		}
	}
	backtrack(0)
	return ans
}

// @lc code=end
