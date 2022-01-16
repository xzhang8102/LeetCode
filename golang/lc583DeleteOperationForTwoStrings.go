package golang

/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] Delete Operation for Two Strings
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	for i := 0; i <= len1; i++ {
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}
	for index1 := 1; index1 <= len1; index1++ {
		for index2 := 1; index2 <= len2; index2++ {
			if word1[index1-1] == word2[index2-1] {
				dp[index1][index2] = dp[index1-1][index2-1]
			} else {
				dp[index1][index2] = 1 + lc583Min(dp[index1-1][index2], dp[index1][index2-1])
			}
		}
	}
	return dp[len1][len2]
}

func lc583Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
