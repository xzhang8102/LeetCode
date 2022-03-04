package golang

import "strings"

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] Integer to Roman
 */

// @lc code=start
func intToRoman(num int) string {
	bases := []struct {
		val   int
		roman string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"}, {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}
	var b strings.Builder
	for _, base := range bases {
		for num >= base.val {
			b.WriteString(base.roman)
			num -= base.val
		}
		if num == 0 {
			break
		}
	}
	return b.String()
}

// @lc code=end
