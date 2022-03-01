package golang

import "strings"

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	n := len(s)
	if numRows <= 1 || numRows >= n {
		return s
	}
	cache := make([][]byte, numRows)
	for i := range cache {
		cache[i] = make([]byte, 0)
	}
	for i, dir, rowN := 0, -1, 0; i < n; i, rowN = i+1, rowN+dir {
		cache[rowN] = append(cache[rowN], s[i])
		if rowN == numRows-1 || rowN == 0 {
			dir *= -1
		}
	}
	var b strings.Builder
	b.Grow(n)
	for _, row := range cache {
		b.Write(row)
	}
	return b.String()
}

// @lc code=end
