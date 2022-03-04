package golang

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 1 {
		return strs[0]
	}
	sort.SliceStable(strs, func(i, j int) bool { return len(strs[i]) < len(strs[j]) })
	var b strings.Builder
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		valid := true
		for j := 1; j < n; j++ {
			if strs[j][i] != char {
				valid = false
				break
			}
		}
		if valid {
			b.WriteByte(char)
		} else {
			break
		}
	}
	return b.String()
}

// @lc code=end
