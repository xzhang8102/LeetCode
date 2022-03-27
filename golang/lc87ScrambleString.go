package golang

/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] Scramble String
 */

// @lc code=start
func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if !lc87Check(s1, s2) {
		return false
	}
	n := len(s1)
	for i := 1; i < n; i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i]) {
			return true
		}
	}
	return false
}

func lc87Check(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	dict := [26]int{}
	for i := range s1 {
		dict[s1[i]-'a']++
		dict[s2[i]-'a']--
	}
	for _, v := range dict {
		if v != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
