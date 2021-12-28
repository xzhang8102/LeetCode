package golang

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// cover the case when string contains unicode
	// if the input only consists of ASCII characters
	// could use a 26-byte array instead
	dict := make(map[rune]int)
	for _, r := range s {
		dict[r]++
	}
	for _, r := range t {
		count, ok := dict[r]
		if !ok {
			return false
		}
		if count == 0 {
			return false
		}
		dict[r]--
	}
	return true
}

// @lc code=end
