package golang

import "strings"

/*
 * @lc app=leetcode.cn id=804 lang=golang
 *
 * [804] Unique Morse Code Words
 */

// @lc code=start
func uniqueMorseRepresentations(words []string) int {
	mapping := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	dict := map[string]bool{}
	var b strings.Builder
	for _, word := range words {
		for _, ch := range word {
			b.WriteString(mapping[ch-'a'])
		}
		dict[b.String()] = true
		b.Reset()
	}
	return len(dict)
}

// @lc code=end
