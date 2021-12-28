package golang

/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	var result [][]int
	var backtrack func(num int, combination []int)
	backtrack = func(num int, combination []int) {
		combination = append(combination, num)
		if k == len(combination) {
			tmp := make([]int, len(combination))
			copy(tmp, combination)
			result = append(result, tmp)
			return
		}
		for i := num + 1; i <= n; i++ {
			backtrack(i, combination)
		}
	}
	for i := 1; i <= n; i++ {
		backtrack(i, []int{})
	}
	return result
}

// @lc code=end
