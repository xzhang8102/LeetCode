package golang

/*
 * @lc app=leetcode.cn id=953 lang=golang
 *
 * [953] Verifying an Alien Dictionary
 */

// @lc code=start
func isAlienSorted(words []string, order string) bool {
	dict := make(map[rune]int, 26)
	for i, ch := range order {
		dict[ch] = i
	}
	for i := 1; i < len(words); i++ {
		if lc953Compare(words[i-1], words[i], dict) {
			return false
		}
	}
	return true
}

func lc953Compare(s1, s2 string, dict map[rune]int) bool {
	i, j := 0, 0
	for ; i < len(s1) && j < len(s2); i, j = i+1, j+1 {
		if dict[rune(s1[i])] > dict[rune(s2[j])] {
			return true
		}
		if dict[rune(s1[i])] < dict[rune(s2[j])] {
			return false
		}
	}
	if i < len(s1) {
		return true
	}
	return false
}

// @lc code=end
