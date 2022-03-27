package weekly286

import (
	"math"
)

func kthPalindrome(queries []int, intLength int) []int64 {
	ans := []int64{}
	// max counts of palindrome number starting from 1 ~ 9
	max := int(math.Pow10((intLength - 1) / 2))
	for _, query := range queries {
		if query > 9*max {
			ans = append(ans, -1)
		} else {
			n := int64(query - 1 + max)
			i := n
			if intLength&1 == 1 {
				i /= 10
			}
			for i > 0 {
				n = n*10 + i%10
				i /= 10
			}
			ans = append(ans, n)
		}
	}
	return ans
}
