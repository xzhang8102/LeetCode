package golang

import (
	"unicode"
)

/*
 * @lc app=leetcode.cn id=1935 lang=golang
 *
 * [1935] Maximum Number of Words You Can Type
 */

// @lc code=start
func canBeTypedWords(text string, brokenLetters string) int {
	var count int = 0
	set := map[rune]struct{}{}
	for _, c := range brokenLetters {
		set[c] = struct{}{}
	}
	affected := false
	for _, char := range text {
		if unicode.IsSpace(char) {
			if !affected {
				count++
			}
			affected = false
		} else {
			if !affected {
				_, ok := set[char]
				if ok {
					affected = true
				}
			}
		}
	}
	// Don't forget to check status when loop to the end
	if !affected {
		count++
	}
	return count
}

// @lc code=end
