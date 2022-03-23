package golang

/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

// @lc code=start
func convertToTitle(columnNumber int) string {
	ans := []byte{}
	for columnNumber > 0 {
		columnNumber--
		ans = append(ans, byte('A'+columnNumber%26))
		columnNumber /= 26
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}
	return string(ans)
}

// @lc code=end
