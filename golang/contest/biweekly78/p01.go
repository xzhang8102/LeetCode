package biweekly78

import "math"

func divisorSubstrings(num int, k int) int {
	base := int(math.Pow10(k))
	lower := int(math.Pow10(k - 1))
	ans := 0
	origin := num
	for num >= lower {
		n := num % base
		if n != 0 && origin%n == 0 {
			ans++
		}
		num /= 10
	}
	return ans
}
