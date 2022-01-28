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
		return properties[i][0] > properties[j][0] || (properties[i][0] == properties[j][0] && properties[i][1] < properties[j][1])
	})
	ans := 0
	maxDef := 0
	for _, prop := range properties {
		if prop[1] < maxDef {
			ans++
		} else {
			maxDef = prop[1]
		}
	}
	return ans
}

// @lc code=end
