package golang

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start
func numDecodings(s string) int {
	n := len(s)
	prev1, prev2 := 1, 0
	for i := 1; i <= n; i++ {
		tmp := 0
		if s[i-1] != '0' {
			tmp += prev1
		}
		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
			tmp += prev2
		}
		prev1, prev2 = tmp, prev1
	}
	return prev1
}

// @lc code=end
