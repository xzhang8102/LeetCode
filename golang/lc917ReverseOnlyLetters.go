package golang

import "unicode"

/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] Reverse Only Letters
 */

// @lc code=start
func reverseOnlyLetters(s string) string {
	n := len(s)
	ans := []byte(s)
	for left, right := 0, n-1; left < right; {
		if !unicode.IsLetter(rune(ans[left])) {
			left++
			continue
		}
		if !unicode.IsLetter(rune(ans[right])) {
			right--
			continue
		}
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}
	return string(ans)
}

// @lc code=end
