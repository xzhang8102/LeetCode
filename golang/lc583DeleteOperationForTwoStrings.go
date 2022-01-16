package golang

/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] Delete Operation for Two Strings
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	memo := make([][]int, len1)
	for i := range memo {
		memo[i] = make([]int, len2)
		for k := range memo[i] {
			memo[i][k] = -1
		}
	}
	var dp func(int, int) int
	dp = func(index1, index2 int) int {
		if index1 == len1 {
			return len2 - index2
		}
		if index2 == len2 {
			return len1 - index1
		}
		if memo[index1][index2] != -1 {
			return memo[index1][index2]
		}
		if word1[index1] != word2[index2] {
			memo[index1][index2] = 1 + lc583Min(dp(index1+1, index2), dp(index1, index2+1))
		} else {
			memo[index1][index2] = dp(index1+1, index2+1)
		}
		return memo[index1][index2]
	}
	return dp(0, 0)
}

func lc583Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
