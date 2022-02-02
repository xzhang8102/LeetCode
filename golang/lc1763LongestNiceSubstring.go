package golang

import "unicode"

/*
 * @lc app=leetcode.cn id=1763 lang=golang
 *
 * [1763] Longest Nice Substring
 */

// @lc code=start
func longestNiceSubstring(s string) string {
	lower, upper := 0, 0
	for _, ch := range s {
		if unicode.IsLower(ch) {
			lower |= 1 << (ch - 'a')
		} else {
			upper |= 1 << (ch - 'A')
		}
	}
	if lower == upper {
		return s
	}
	valid := lower & upper
	ans := ""
	for i := 0; i < len(s); i++ {
		start := i
		for i < len(s) && valid>>(unicode.ToLower(rune(s[i]))-'a')&1 == 1 {
			i++
		}
		if t := longestNiceSubstring(s[start:i]); len(t) > len(ans) {
			ans = t
		}
	}
	return ans
}

// @lc code=end
