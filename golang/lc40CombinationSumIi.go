package golang

import "sort"

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	ans := [][]int{}
	pick := []int{}
	sort.Ints(candidates)
	visited := make([]bool, len(candidates))
	var backtrack func(int, int)
	backtrack = func(index, remain int) {
		if index > len(candidates) {
			return
		}
		if remain == 0 {
			ans = append(ans, append([]int(nil), pick...))
			return
		}
		for i := index; i < len(candidates); i++ {
			if remain-candidates[i] >= 0 {
				if i == 0 || (i > 0 && candidates[i-1] != candidates[i]) || (candidates[i-1] == candidates[i] && visited[i-1]) {
					visited[i] = true
					pick = append(pick, candidates[i])
					backtrack(i+1, remain-candidates[i])
					pick = pick[:len(pick)-1]
					visited[i] = false
				}
			}
		}
	}
	backtrack(0, target)
	return ans
}

// @lc code=end
