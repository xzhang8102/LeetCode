package golang

/*
 * @lc app=leetcode.cn id=829 lang=golang
 *
 * [829] Consecutive Numbers Sum
 */

// @lc code=start
func consecutiveNumbersSum(n int) int {
	// assume `a` is the first element of the `k` consecutive numbers of which the sum is `n`
	// (a + a + k - 1) k / 2 = n  -->  2a = 2n / k - k + 1
	// because a and k both are greater than 1
	// 2a >= 2 --> 2n / k - k + 1 >= 2 --> 2n / k >= k + 1 --> 2n / k > k
	ans := 0
	for k := 1; k*k < 2*n; k++ {
		// (a + a + k - 1) k / 2 = n
		// (2a + k - 1) * k = 2n
		if 2*n%k != 0 {
			continue
		}
		// 2a = 2n / k - k + 1
		if (2*n/k-k+1)&1 == 0 {
			ans++
		}
	}
	return ans
}

// @lc code=end
