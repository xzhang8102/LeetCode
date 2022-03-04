package golang

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	if len(strs) == 2 {
		str1, str2 := strs[0], strs[1]
		n := lc14Min(len(str1), len(str2))
		i := 0
		for ; i < n && str1[i] == str2[i]; i++ {
		}
		return str1[:i]
	}
	mid := len(strs) >> 1
	return longestCommonPrefix([]string{longestCommonPrefix(strs[:mid]), longestCommonPrefix(strs[mid:])})
}

func lc14Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
