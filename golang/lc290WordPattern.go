package golang

import "strings"

/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] Word Pattern
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	mappingP := [26]string{}
	mappingS := map[string]rune{}
	for i, ch := range pattern {
		if mappingP[ch-'a'] == "" {
			mappingP[ch-'a'] = words[i]
		} else if mappingP[ch-'a'] != words[i] {
			return false
		}
		if mappingS[words[i]] == 0 {
			mappingS[words[i]] = ch
		} else if mappingS[words[i]] != ch {
			return false
		}
	}
	return true
}

// @lc code=end
