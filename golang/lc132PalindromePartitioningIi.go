package golang

/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] Palindrome Partitioning II
 */

// @lc code=start
func minCut(s string) int {
	n := len(s)
	check := make([][]bool, n)
	for i := range check {
		check[i] = make([]bool, n)
		check[i][i] = true
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			check[i][j] = s[i] == s[j]
			if j > i+2 {
				check[i][j] = check[i][j] && check[i+1][j-1]
			}
		}
	}
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if check[0][i] {
			dp[i] = 0
		} else {
			dp[i] = i
			for j := 1; j <= i; j++ {
				if check[j][i] {
					dp[i] = lc132Min(dp[i], dp[j-1]+1)
				}
			}
		}
	}
	return dp[n-1]
}

func lc132Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
