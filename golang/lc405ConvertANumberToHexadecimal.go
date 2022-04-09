package golang

import "strings"

/*
 * @lc app=leetcode.cn id=405 lang=golang
 *
 * [405] Convert a Number to Hexadecimal
 */

// @lc code=start
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var b strings.Builder
	for i := 7; i >= 0; i-- {
		n := (num >> (4 * i)) & 0x0f
		if n != 0 || b.Len() > 0 {
			switch {
			case n >= 10:
				b.WriteByte(byte(n-10) + 'a')
			default:
				b.WriteByte(byte(n) + '0')
			}
		}
	}
	return b.String()
}

// @lc code=end
