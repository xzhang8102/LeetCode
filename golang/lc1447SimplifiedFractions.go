package golang

import "fmt"

/*
 * @lc app=leetcode.cn id=1447 lang=golang
 *
 * [1447] Simplified Fractions
 */

// @lc code=start
func simplifiedFractions(n int) []string {
	ans := []string{}
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if lc1447gcd(i, j) == 1 {
				ans = append(ans, fmt.Sprintf("%d/%d", j, i))
			}
		}
	}
	return ans
}

func lc1447gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// @lc code=end
