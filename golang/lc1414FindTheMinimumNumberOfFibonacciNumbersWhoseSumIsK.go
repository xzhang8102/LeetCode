package golang

/*
 * @lc app=leetcode.cn id=1414 lang=golang
 *
 * [1414] Find the Minimum Number of Fibonacci Numbers Whose Sum Is K
 */

// @lc code=start
func findMinFibonacciNumbers(k int) int {
	fib := []int{}
	for i, j := 1, 1; i <= k; i, j = i+j, i {
		fib = append(fib, i)
	}
	ans := 0
	for i := len(fib) - 1; k > 0; i-- {
		if k >= fib[i] {
			ans++
			k -= fib[i]
		}
	}
	return ans
}

// @lc code=end
