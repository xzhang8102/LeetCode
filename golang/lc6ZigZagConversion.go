package golang

import "strings"

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	buffer := make([][]byte, 0)
	for i := 0; i < numRows; i++ {
		buffer = append(buffer, []byte{})
	}
	n := len(s)
	for i, row, dir := 0, 0, 1; i < n; i++ {
		buffer[row] = append(buffer[row], s[i])
		if row == 0 {
			dir = 1
		}
		if row == numRows-1 {
			dir = -1
		}
		row += dir
	}
	var b strings.Builder
	b.Grow(n)
	for _, r := range buffer {
		b.Write(r)
	}
	return b.String()
}

// @lc code=end
