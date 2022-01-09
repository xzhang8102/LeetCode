package golang

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	pick := []int{}
	n := len(candidates)
	var backtrack func(int, int)
	backtrack = func(index, target int) {
		if index == n {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), pick...))
			return
		}
		backtrack(index+1, target)
		if target-candidates[index] >= 0 {
			pick = append(pick, candidates[index])
			backtrack(index, target-candidates[index])
			pick = pick[:len(pick)-1]
		}
	}
	backtrack(0, target)
	return ans
}

// @lc code=end
