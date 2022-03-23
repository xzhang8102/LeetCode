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
		char := columnNumber % 26
		columnNumber /= 26
		if char == 0 {
			ans = append(ans, 'Z')
			columnNumber--
		} else {
			ans = append(ans, byte('A'+char-1))
		}
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}

// @lc code=end
