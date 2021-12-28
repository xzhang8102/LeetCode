package golang

/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	len1, len2 := len(s1), len(s2)
	charCount := [26]byte{}
	for i := 0; i < len1; i++ {
		charCount[s1[i]-'a']++
	}
	l, r := 0, 0
	for l <= len2-len1 {
		for r < l+len1 && charCount[s2[r]-'a'] >= 1 {
			charCount[s2[r]-'a']--
			r++
		}
		if r == l+len1 {
			return true
		}
		charCount[s2[l]-'a']++
		l++
	}
	return false
}

// @lc code=end
