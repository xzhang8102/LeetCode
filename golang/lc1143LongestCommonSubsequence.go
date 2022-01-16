package golang

/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	for index1 := 1; index1 <= len1; index1++ {
		for index2 := 1; index2 <= len2; index2++ {
			if text1[index1-1] == text2[index2-1] {
				dp[index1][index2] = 1 + dp[index1-1][index2-1]
			} else {
				dp[index1][index2] = lc1143Max(dp[index1-1][index2], dp[index1][index2-1])
			}
		}
	}
	return dp[len1][len2]
}

func lc1143Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
