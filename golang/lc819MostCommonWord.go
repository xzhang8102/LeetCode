package golang

import (
	"strings"
	"unicode"
)

/*
 * @lc app=leetcode.cn id=819 lang=golang
 *
 * [819] Most Common Word
 */

// @lc code=start
func mostCommonWord(paragraph string, banned []string) string {
	dict := map[string]int{}
	for _, w := range banned {
		dict[w] = -1
	}
	max := 0
	ans := ""
	start := 0
	n := len(paragraph)
	for i := 0; i <= n; i++ {
		if i == n || !unicode.IsLetter(rune(paragraph[i])) {
			if i == start {
				start++
			} else {
				if start < n {
					w := strings.ToLower(paragraph[start:i])
					if dict[w] != -1 {
						dict[w]++
						if dict[w] > max {
							max = dict[w]
							ans = w
						}
					}
					start = i + 1
				}
			}
		}
	}
	return ans
}

// @lc code=end
