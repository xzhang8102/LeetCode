package golang

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	ans := "1"
	for n > 1 {
		var b strings.Builder
		for i := 0; i < len(ans); {
			j := i + 1
			for ; j < len(ans) && ans[j] == ans[i]; j++ {
			}
			b.WriteString(strconv.Itoa(j - i))
			b.WriteByte(ans[i])
			i = j
		}
		ans = b.String()
		n--
	}
	return ans
}

// @lc code=end
