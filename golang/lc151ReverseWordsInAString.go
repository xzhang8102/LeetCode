package golang

import "strings"

/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	ans := make([]string, 0, len(arr))
	for _, word := range arr {
		if word != "" {
			ans = append(ans, word)
		}
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return strings.Join(ans, " ")
}

// @lc code=end
