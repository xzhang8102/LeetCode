package golang

import (
	"sort"
	"strings"
	"unicode"
)

/*
 * @lc app=leetcode.cn id=937 lang=golang
 *
 * [937] Reorder Data in Log Files
 */

// @lc code=start
func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		s1 := strings.SplitN(logs[i], " ", 2)[1]
		s2 := strings.SplitN(logs[j], " ", 2)[1]
		isDigit1 := unicode.IsDigit(rune(s1[0]))
		isDigit2 := unicode.IsDigit(rune(s2[0]))
		if isDigit1 && isDigit2 {
			return false
		}
		if !isDigit1 && !isDigit2 {
			return s1 < s2 || s1 == s2 && logs[i] < logs[j]
		}
		return !isDigit1
	})
	return logs
}

// @lc code=end
