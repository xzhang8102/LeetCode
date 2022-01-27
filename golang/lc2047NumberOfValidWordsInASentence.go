package golang

import "strings"

/*
 * @lc app=leetcode.cn id=2047 lang=golang
 *
 * [2047] Number of Valid Words in a Sentence
 */

// @lc code=start
func countValidWords(sentence string) int {
	tokens := strings.Fields(sentence)
	ans := 0
	for _, token := range tokens {
		if lc2047IsValid(token) {
			ans++
		}
	}
	return ans
}

func lc2047IsValid(token string) bool {
	n := len(token)
	hasHyphen := false
	for i := 0; i < n; i++ {
		char := token[i]
		if char >= '0' && char <= '9' {
			return false
		}
		switch char {
		case '-':
			if i == 0 || i == n-1 || hasHyphen || token[i+1] < 'a' || token[i+1] > 'z' {
				return false
			}
			hasHyphen = true
		case '!', '.', ',':
			if i != n-1 {
				return false
			}
		}
	}
	return true
}

// @lc code=end
