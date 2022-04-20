package golang

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=824 lang=golang
 *
 * [824] Goat Latin
 */

// @lc code=start
func toGoatLatin(sentence string) string {
	words := strings.Split(sentence, " ")
	for i, word := range words {
		convert := ""
		switch word[0] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			convert = fmt.Sprintf("%sma", word)
		default:
			convert = fmt.Sprintf("%s%cma", word[1:], word[0])
		}
		convert = fmt.Sprintf("%s%s", convert, strings.Repeat("a", i+1))
		words[i] = convert
	}
	return strings.Join(words, " ")
}

// @lc code=end
