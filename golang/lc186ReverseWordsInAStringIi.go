package golang

/*
 * @lc app=leetcode.cn id=186 lang=golang
 *
 * [186] Reverse Words in a String II
 */

// @lc code=start
func reverseWords(s []byte) {
	n := len(s)
	lc186Reverse(s, 0, n-1)
	i := 0
	for i < n {
		j := i
		for ; j < n && s[j] != ' '; j++ {
		}
		lc186Reverse(s, i, j-1)
		i = j + 1
	}
}

func lc186Reverse(s []byte, start, end int) {
	for ; start < end; start, end = start+1, end-1 {
		s[start], s[end] = s[end], s[start]
	}
}

// @lc code=end
