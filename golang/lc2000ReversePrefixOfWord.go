package golang

import "strings"

/*
 * @lc app=leetcode.cn id=2000 lang=golang
 *
 * [2000] Reverse Prefix of Word
 */

// @lc code=start
func reversePrefix(word string, ch byte) string {
	right := strings.IndexByte(word, ch)
	if right < 0 {
		return word
	}
	s := []byte(word)
	for left := 0; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
	return string(s)
}

// @lc code=end
