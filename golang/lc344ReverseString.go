package golang

/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] Reverse String
 */

// @lc code=start
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// @lc code=end
