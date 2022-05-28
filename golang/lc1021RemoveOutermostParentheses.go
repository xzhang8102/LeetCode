package golang

import "strings"

/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] Remove Outermost Parentheses
 */

// @lc code=start
func removeOuterParentheses(s string) string {
	state := 0
	start := 0
	var b strings.Builder
	for i, ch := range s {
		if ch == '(' {
			if state == 0 {
				start = i
			}
			state++
		}
		if ch == ')' {
			state--
			if state == 0 {
				b.WriteString(s[start+1 : i])
			}
		}
	}
	return b.String()
}

// @lc code=end
