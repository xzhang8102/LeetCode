package golang

import "sort"

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	ans := [][]int{}
	pick := []int{}
	n := len(nums)
	sort.Ints(nums)
	var backtrack func(index int)
	backtrack = func(index int) {
		tmp := []int{}
		tmp = append(tmp, pick...)
		ans = append(ans, tmp)
		for i := index; i < n; i++ {
			if i == index || nums[i] != nums[i-1] {
				pick = append(pick, nums[i])
				backtrack(i + 1)
				pick = pick[:len(pick)-1]
			}
		}
	}
	backtrack(0)
	return ans
}

// @lc code=end
