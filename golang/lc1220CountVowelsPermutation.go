package golang

/*
 * @lc app=leetcode.cn id=1220 lang=golang
 *
 * [1220] Count Vowels Permutation
 */

// @lc code=start
func countVowelPermutation(n int) int {
	const mod int = 1e9 + 7
	dp := [5]int{1, 1, 1, 1, 1}
	for L := 1; L < n; L++ {
		// a - 0 e - 1 i - 2 o - 3 u - 4
		// a - e
		// e - a/i
		// i - a/e/o/u
		// o - i/u
		// u - a
		dp = [5]int{
			(dp[1] + dp[2] + dp[4]) % mod,
			(dp[0] + dp[2]) % mod,
			(dp[1] + dp[3]) % mod,
			dp[2],
			(dp[2] + dp[3]) % mod,
		}
	}
	ans := 0
	for _, v := range dp {
		ans = (ans + v) % mod
	}
	return ans
}

// @lc code=end
