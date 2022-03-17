package golang

import "sort"

/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] Longest Word in Dictionary
 */

// @lc code=start
func longestWord(words []string) string {
	sort.SliceStable(words, func(i, j int) bool {
		s, t := words[i], words[j]
		// make sure to process the word with smaller lexicographical order later
		return len(s) < len(t) || len(s) == len(t) && s > t
	})
	candidates := map[string]bool{"": true}
	longest := ""
	for _, word := range words {
		if candidates[word[:len(word)-1]] {
			longest = word
			candidates[word] = true
		}
	}
	return longest
}

// @lc code=end
