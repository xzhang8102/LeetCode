package golang

import "strings"

/*
 * @lc app=leetcode.cn id=884 lang=golang
 *
 * [884] Uncommon Words from Two Sentences
 */

// @lc code=start
func uncommonFromSentences(s1 string, s2 string) []string {
	dict := map[string]int{}
	for _, s := range strings.Fields(s1 + " " + s2) {
		dict[s]++
	}
	ans := []string{}
	for s, cnt := range dict {
		if cnt == 1 {
			ans = append(ans, s)
		}
	}
	return ans
}

// @lc code=end
