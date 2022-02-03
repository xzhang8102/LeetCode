package golang

/*
 * @lc app=leetcode.cn id=1414 lang=golang
 *
 * [1414] Find the Minimum Number of Fibonacci Numbers Whose Sum Is K
 */

// @lc code=start
func findMinFibonacciNumbers(k int) int {
	first, second := 1, 1
	for first <= k {
		first, second = first+second, first
	}
	ans := 0
	for k > 0 {
		if k-second >= 0 {
			k -= second
			ans++
		}
		first, second = second, first-second
	}
	return ans
}

// @lc code=end
