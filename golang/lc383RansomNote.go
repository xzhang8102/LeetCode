package golang

/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	dict := map[rune]int{}
	for _, ch := range magazine {
		dict[ch]++
	}
	for _, ch := range ransomNote {
		if dict[ch] > 0 {
			dict[ch]--
		} else {
			return false
		}
	}
	return true
}

// @lc code=end
