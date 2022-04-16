package golang

/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
func partition(s string) [][]string {
	ans := [][]string{}
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j]
			if j-i > 1 && dp[i][j] {
				dp[i][j] = dp[i][j] && dp[i+1][j-1]
			}
		}
	}
	pick := []string{}
	var dfs func(start int)
	dfs = func(start int) {
		if start == n {
			tmp := make([]string, len(pick))
			copy(tmp, pick)
			ans = append(ans, tmp)
			return
		}
		for i := start; i < n; i++ {
			if dp[start][i] {
				pick = append(pick, s[start:i+1])
				dfs(i + 1)
				pick = pick[:len(pick)-1]
			}
		}
	}
	dfs(0)
	return ans
}

// @lc code=end
