package golang

import "math"

/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] Power of Three
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	return n > 0 && int(math.Pow(3, 19))%n == 0
}

// @lc code=end
