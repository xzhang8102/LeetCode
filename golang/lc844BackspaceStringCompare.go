package golang

/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] Backspace String Compare
 */

// @lc code=start
func backspaceCompare(s string, t string) bool {
	for sp, tp := len(s)-1, len(t)-1; sp >= 0 || tp >= 0; sp, tp = sp-1, tp-1 {
		skips := 0
		for sp >= 0 {
			if s[sp] == '#' {
				skips++
				sp--
			} else if skips > 0 {
				skips--
				sp--
			} else {
				break
			}
		}
		skipt := 0
		for tp >= 0 {
			if t[tp] == '#' {
				skipt++
				tp--
			} else if skipt > 0 {
				skipt--
				tp--
			} else {
				break
			}
		}
		if sp >= 0 && tp >= 0 {
			if s[sp] != t[tp] {
				return false
			}
		} else if sp >= 0 || tp >= 0 {
			return false
		}
	}
	return true
}

// @lc code=end
