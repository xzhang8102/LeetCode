package golang

import "strings"

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] Integer to Roman
 */

// @lc code=start
func intToRoman(num int) string {
	mapping := map[int]rune{
		1:    'I',
		5:    'V',
		10:   'X',
		50:   'L',
		100:  'C',
		500:  'D',
		1000: 'M',
	}
	bases := []int{1000, 500, 100, 50, 10, 5, 1}
	var b strings.Builder
	for _, base := range bases {
		if quotient := num / base; quotient >= 1 {
			switch {
			case base == 500 && num >= 900:
				b.WriteString("CM")
				num -= 900
			case base == 100 && num >= 400:
				b.WriteString("CD")
				num -= 400
			case base == 50 && num >= 90:
				b.WriteString("XC")
				num -= 90
			case base == 10 && num >= 40:
				b.WriteString("XL")
				num -= 40
			case base == 5 && num == 9:
				b.WriteString("IX")
				num -= 9
			case base == 1 && num == 4:
				b.WriteString("IV")
				num -= 4
			default:
				b.WriteString(strings.Repeat(string(mapping[base]), quotient))
				num %= base
			}
		}
		if num == 0 {
			break
		}
	}
	return b.String()
}

// @lc code=end
