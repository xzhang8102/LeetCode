package golang

/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */

// @lc code=start
func lc557reverseWords(s string) string {
	n := len(s)
	ans := make([]byte, 0)
	for i := 0; i < n; {
		start := i
		for i < n && s[i] != ' ' {
			i++
		}
		// i -> ' '
		for p := i - 1; p >= start; p-- {
			ans = append(ans, s[p])
		}
		for i < n && s[i] == ' ' {
			ans = append(ans, ' ')
			i++
		}
	}
	return string(ans)
}

// @lc code=end
