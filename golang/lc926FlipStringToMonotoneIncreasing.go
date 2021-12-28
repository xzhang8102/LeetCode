package golang

/*
 * @lc app=leetcode.cn id=926 lang=golang
 *
 * [926] Flip String to Monotone Increasing
 */

// @lc code=start
func minFlipsMonoIncr(s string) int {
	n := len(s)
	// dp[i][1] index=i字符为1时需要flip的次数
	// dp[i][0] index=i字符为0时需要flip的次数
	dp := make([][2]int, n)
	dp[0][0] = int(s[0] - '0')
	dp[0][1] = int(s[0] - '1')
	for i := 1; i < n; i++ {
		// 当前字符为0前一状态必不可能是1
		dp[i][0] = dp[i-1][0] + int(s[i]-'0')
		dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + int('1'-s[i])
	}
	return min(dp[n-1][0], dp[n-1][1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
