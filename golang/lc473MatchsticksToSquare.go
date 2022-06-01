package golang

import "sort"

/*
 * @lc app=leetcode.cn id=473 lang=golang
 *
 * [473] Matchsticks to Square
 */

// @lc code=start
func makesquare(matchsticks []int) bool {
	n := len(matchsticks)
	if n < 4 {
		return false
	}
	totalLen := 0
	for _, match := range matchsticks {
		totalLen += match
	}
	if totalLen%4 != 0 {
		return false
	}
	edge := totalLen / 4
	sort.Slice(matchsticks, func(i, j int) bool { return matchsticks[i] >= matchsticks[j] })
	found := false
	pick := [4]int{}
	visited := map[[4]int]bool{}
	var backtrack func(start int)
	backtrack = func(start int) {
		if found {
			return
		}
		if visited[pick] {
			return
		}
		visited[pick] = true
		if start == n {
			found = true
			return
		}
		for i := 0; i < 4; i++ {
			if pick[i]+matchsticks[start] <= edge {
				pick[i] += matchsticks[start]
				backtrack(start + 1)
				pick[i] -= matchsticks[start]
			}
		}
	}
	backtrack(0)
	return found
}

// @lc code=end
