package golang

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=408 lang=golang
 *
 * [408] Valid Word Abbreviation
 */

// @lc code=start
func validWordAbbreviation(word string, abbr string) bool {
	n, m := len(word), len(abbr)
	var b strings.Builder
	b.Grow(n)
	i, j := 0, 0
	for i < m {
		if abbr[i] >= 'a' && abbr[i] <= 'z' {
			b.WriteByte(abbr[i])
			i++
			j++
		} else {
			if abbr[i] == '0' {
				return false
			}
			start := i
			for i < m && abbr[i] >= '0' && abbr[i] <= '9' {
				i++
			}
			if i > start {
				numLen, _ := strconv.Atoi(abbr[start:i])
				if j+numLen > n {
					return false
				}
				b.WriteString(word[j : j+numLen])
				j += numLen
			}
		}
	}
	return b.String() == word
}

// @lc code=end
