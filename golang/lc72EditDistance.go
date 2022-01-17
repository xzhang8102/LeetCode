package golang

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	memo := make([][]int, len1)
	for i := range memo {
		memo[i] = make([]int, len2)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dp func(int, int) int
	dp = func(i1, i2 int) int {
		if i1 == len1 {
			return len2 - i2
		}
		if i2 == len2 {
			return len1 - i1
		}
		if memo[i1][i2] != -1 {
			return memo[i1][i2]
		}
		if word1[i1] == word2[i2] {
			memo[i1][i2] = dp(i1+1, i2+1)
		} else {
			memo[i1][i2] = lc72Min(
				1+dp(i1+1, i2+1),
				1+dp(i1+1, i2),
				1+dp(i1, i2+1),
			)
		}
		return memo[i1][i2]
	}
	return dp(0, 0)
}

func lc72Min(nums ...int) int {
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}

// @lc code=end
