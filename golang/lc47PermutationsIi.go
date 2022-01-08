package golang

import "sort"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	pick := []int{}
	n := len(nums)
	visited := make([]bool, n)
	var backtrack func(int)
	backtrack = func(count int) {
		if count == n {
			ans = append(ans, append([]int(nil), pick...))
			return
		}
		for i, v := range nums {
			if visited[i] || i > 0 && !visited[i-1] && v == nums[i-1] {
				continue
			}
			visited[i] = true
			pick = append(pick, v)
			backtrack(count + 1)
			visited[i] = false
			pick = pick[:len(pick)-1]
		}
	}
	backtrack(0)
	return ans
}

// @lc code=end
