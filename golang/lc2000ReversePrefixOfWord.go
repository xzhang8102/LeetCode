package golang

/*
 * @lc app=leetcode.cn id=2000 lang=golang
 *
 * [2000] Reverse Prefix of Word
 */

// @lc code=start
func reversePrefix(word string, ch byte) string {
	n := len(word)
	buf := make([]byte, n)
	index := -1
	for i := 0; i < n; i++ {
		buf[n-i-1] = word[i]
		if word[i] == ch {
			index = i
			break
		}
	}
	if index == -1 {
		return word
	} else {
		return string(buf[n-1-index:]) + word[index+1:]
	}
}

// @lc code=end
