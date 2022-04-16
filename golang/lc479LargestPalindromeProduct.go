package golang

import "math"

/*
 * @lc app=leetcode.cn id=479 lang=golang
 *
 * [479] Largest Palindrome Product
 */

// @lc code=start
func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	max := int(math.Pow10(n)) - 1
	for i := max; i >= 0; i-- {
		num, j := i, i
		for j > 0 {
			num = num*10 + j%10
			j /= 10
		}
		for k := max; k*k >= num; k-- {
			if num%k == 0 {
				return num % 1337
			}
		}
	}
	return -1
}

// @lc code=end
