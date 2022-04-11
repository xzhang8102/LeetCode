package golang

/*
 * @lc app=leetcode.cn id=408 lang=golang
 *
 * [408] Valid Word Abbreviation
 */

// @lc code=start
func validWordAbbreviation(word string, abbr string) bool {
	n, m := len(word), len(abbr)
	i, j := 0, 0
	for i < m && j < n {
		if abbr[i] >= 'a' && abbr[i] <= 'z' {
			if abbr[i] != word[j] {
				return false
			}
			i++
			j++
		} else {
			if abbr[i] == '0' {
				return false
			}
			numLen := 0
			for i < m && abbr[i] >= '0' && abbr[i] <= '9' {
				numLen = numLen*10 + int(abbr[i]-'0')
				i++
			}
			j += numLen
		}
	}
	return i == m && j == n
}

// @lc code=end
