package golang

import "sort"

/*
 * @lc app=leetcode.cn id=1996 lang=golang
 *
 * [1996] The Number of Weak Characters in the Game
 */

// @lc code=start
func numberOfWeakCharacters(properties [][]int) int {
	sort.SliceStable(properties, func(i, j int) bool {
		return properties[i][0] < properties[j][0] || (properties[i][0] == properties[j][0] && properties[i][1] > properties[j][1])
	})
	ans := 0
	stack := []int{}
	for _, p := range properties {
		for len(stack) > 0 && stack[len(stack)-1] < p[1] {
			stack = stack[:len(stack)-1]
			ans++
		}
		stack = append(stack, p[1])
	}
	return ans
}

// @lc code=end
