package golang

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}
	for i1 := 1; i1 <= len1; i1++ {
		for i2 := 1; i2 <= len2; i2++ {
			if word1[i1-1] == word2[i2-1] {
				dp[i1][i2] = dp[i1-1][i2-1]
			} else {
				dp[i1][i2] = lc72Min(
					// replace current char
					dp[i1-1][i2-1]+1,
					// add the diff char
					lc72Min(
						dp[i1-1][i2]+1,
						dp[i1][i2-1]+1,
					),
				)
			}
		}
	}
	return dp[len1][len2]
}

func lc72Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
