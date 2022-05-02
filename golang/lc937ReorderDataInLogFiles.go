package golang

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=937 lang=golang
 *
 * [937] Reorder Data in Log Files
 */

// @lc code=start
func reorderLogFiles(logs []string) []string {
	n := len(logs)
	ans := make([]string, n)
	letter, digit := 0, n-1
	for _, log := range logs {
		i := strings.IndexByte(log, ' ')
		if log[i+1] >= '0' && log[i+1] <= '9' {
			ans[digit] = log
			digit--
		} else {
			ans[letter] = log
			letter++
		}
	}
	letterLog := ans[:letter]
	sort.Slice(letterLog, func(i, j int) bool {
		s1, s2 := strings.IndexByte(letterLog[i], ' '), strings.IndexByte(letterLog[j], ' ')
		if letterLog[i][s1+1:] != letterLog[j][s2+1:] {
			return letterLog[i][s1+1:] < letterLog[j][s2+1:]
		} else {
			return letterLog[i][:s1] < letterLog[j][:s2]
		}
	})
	for i, j := letter, n-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

// @lc code=end
