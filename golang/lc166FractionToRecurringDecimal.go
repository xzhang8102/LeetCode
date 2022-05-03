package golang

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=166 lang=golang
 *
 * [166] Fraction to Recurring Decimal
 */

// @lc code=start
func fractionToDecimal(numerator int, denominator int) string {
	negative := false
	if numerator < 0 {
		negative = !negative
		numerator = -numerator
	}
	if denominator < 0 {
		negative = !negative
		denominator = -denominator
	}
	integer := numerator / denominator
	numerator -= integer * denominator
	if negative {
		integer = -integer
	}
	if numerator == 0 {
		return strconv.Itoa(integer)
	}
	i := 0
	var b strings.Builder
	pattern := map[int]int{}
	recurring := -1
	for numerator != 0 {
		numerator *= 10
		if index, ok := pattern[numerator]; ok {
			recurring = index
			b.WriteByte(')')
			break
		}
		pattern[numerator] = i
		i++
		quotient := numerator / denominator
		b.WriteByte(byte('0' + quotient))
		numerator -= quotient * denominator
	}
	fraction := b.String()
	var ans string
	if recurring == -1 {
		ans = fmt.Sprintf("%d.%s", integer, fraction)
	} else {
		ans = fmt.Sprintf("%d.%s(%s", integer, fraction[:recurring], fraction[recurring:])
	}
	if integer == 0 && negative {
		ans = "-" + ans
	}
	return ans
}

// @lc code=end
