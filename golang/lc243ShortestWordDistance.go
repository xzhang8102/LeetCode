package golang

import "math"

/*
 * @lc app=leetcode.cn id=243 lang=golang
 *
 * [243] Shortest Word Distance
 */

// @lc code=start
func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}
	ans := math.MaxInt32
	ptr1, ptr2 := -1, -1
	for i, word := range wordsDict {
		if word == word1 {
			ptr1 = i
		}
		if word == word2 {
			ptr2 = i
		}
		if ptr1 >= 0 && ptr2 >= 0 {
			ans = lc243Min(ans, lc243Abs(ptr1-ptr2))
		}
	}
	return ans
}

func lc243Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func lc243Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
