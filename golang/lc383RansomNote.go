package golang

/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	dict := [26]int{}
	for _, ch := range magazine {
		dict[ch-'a']++
	}
	for _, ch := range ransomNote {
		if dict[ch-'a'] > 0 {
			dict[ch-'a']--
		} else {
			return false
		}
	}
	return true
}

// @lc code=end
