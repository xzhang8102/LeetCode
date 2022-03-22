package golang

import "fmt"

/*
 * @lc app=leetcode.cn id=2038 lang=golang
 *
 * [2038] Remove Colored Pieces if Both Neighbors are the Same Color
 */

// @lc code=start
func winnerOfGame(colors string) bool {
	var backtrack func(colors string, turn byte) bool
	backtrack = func(colors string, turn byte) bool {
		n := len(colors)
		for i := 1; i < n-1; i++ {
			if colors[i] == turn && colors[i-1] == turn && colors[i+1] == turn {
				next := fmt.Sprintf("%s%s", colors[:i], colors[i+1:])
				if turn == 'A' {
					return backtrack(next, 'B')
				}
				return backtrack(next, 'A')
			}
		}
		if turn == 'A' {
			return false
		}
		return true
	}
	return backtrack(colors, 'A')
}

// @lc code=end
