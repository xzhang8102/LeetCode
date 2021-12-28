package golang

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	lenNeedle := len(needle)
	if lenNeedle == 0 {
		return 0
	}
	lenHaystack := len(haystack)
	for i := 0; i <= lenHaystack-lenNeedle; i++ {
		if haystack[i:i+lenNeedle] == needle {
			return i
		}
	}
	return -1
}

// @lc code=end
