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
	if n == 1 {
		return "1"
	}
	var ans strings.Builder
	str := countAndSay(n - 1)
	for i := 0; i < len(str); {
		j := i + 1
		for ; j < len(str) && str[j] == str[i]; j++ {
		}
		ans.WriteString(strconv.Itoa(j - i))
		ans.WriteByte(str[i])
		i = j
	}
	return ans.String()
}

// @lc code=end
