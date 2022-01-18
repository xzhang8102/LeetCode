package golang

/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i <= n; i++ {
		max := 0
		for j := i - 1; j >= (i+1)/2; j-- {
			max = lc343Max(max, (i-j)*j, dp[i-j]*dp[j], dp[i-j]*j, (i-j)*dp[j])
		}
		dp[i] = max
	}
	return dp[n]
}

func lc343Max(a ...int) int {
	max := a[0]
	for _, v := range a[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// @lc code=end
