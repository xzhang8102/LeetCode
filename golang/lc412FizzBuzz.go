package golang

import "strconv"

/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := range ans {
		switch {
		case (i+1)%15 == 0:
			ans[i] = "FizzBuzz"
		case (i+1)%3 == 0:
			ans[i] = "Fizz"
		case (i+1)%5 == 0:
			ans[i] = "Buzz"
		default:
			ans[i] = strconv.Itoa(i + 1)
		}
	}
	return ans
}

// @lc code=end
