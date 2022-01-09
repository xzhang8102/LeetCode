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
	var backtrack func(int, int)
	backtrack = func(begin, sum int) {
		if sum == target {
			ans = append(ans, append([]int(nil), pick...))
			return
		}
		for i := begin; i < len(candidates); i++ {
			if i > begin && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i]+sum <= target {
				pick = append(pick, candidates[i])
				backtrack(i+1, candidates[i]+sum)
				pick = pick[:len(pick)-1]
			}
		}
	}
	backtrack(0, 0)
	return ans
}

// @lc code=end
