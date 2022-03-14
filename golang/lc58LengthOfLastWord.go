package golang

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) (ans int) {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	for index >= 0 && s[index] != ' ' {
		ans++
		index--
	}
	return
}

// @lc code=end
