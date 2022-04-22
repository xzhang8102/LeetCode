package golang

/*
 * @lc app=leetcode.cn id=161 lang=golang
 *
 * [161] One Edit Distance
 */

// @lc code=start
func isOneEditDistance(s string, t string) bool {
	n, m := len(s), len(t)
	if n > m {
		return isOneEditDistance(t, s)
	}
	if m-n > 1 {
		return false
	}
	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			if n == m {
				return s[i+1:] == t[i+1:]
			} else {
				return s[i:] == t[i+1:]
			}
		}
	}
	return m-n == 1
}

func lc161Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
