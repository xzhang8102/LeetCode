package golang

import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=564 lang=golang
 *
 * [564] Find the Closest Palindrome
 */

// @lc code=start
func nearestPalindromic(n string) string {
	m := len(n)
	candidates := []int{int(math.Pow10(m-1)) - 1, int(math.Pow10(m)) + 1}
	selfPrefix, _ := strconv.Atoi(n[:(m+1)/2])
	for _, x := range []int{selfPrefix, selfPrefix + 1, selfPrefix - 1} {
		y := x
		if m&1 == 1 {
			y /= 10
		}
		for ; y > 0; y /= 10 {
			x = x*10 + y%10
		}
		candidates = append(candidates, x)
	}
	num, _ := strconv.Atoi(n)
	ans := -1
	for _, candidate := range candidates {
		gap := lc564Abs(candidate - num)
		if gap > 0 {
			if gap < lc564Abs(ans-num) || (gap == lc564Abs(ans-num) && candidate < ans) {
				ans = candidate
			}
		}
	}
	return strconv.Itoa(ans)
}

func lc564Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
