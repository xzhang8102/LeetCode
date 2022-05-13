package golang

/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] Count Primes
 */

// @lc code=start
func countPrimes(n int) int {
	ans := 0
	for i := 2; i < n; i++ {
		if lc204IsPrime(i) {
			ans++
		}
	}
	return ans
}

func lc204IsPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// @lc code=end
