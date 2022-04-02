package golang

import "fmt"

/*
 * @lc app=leetcode.cn id=293 lang=golang
 *
 * [293] Flip Game
 */

// @lc code=start
func generatePossibleNextMoves(currentState string) []string {
	n := len(currentState)
	if n < 2 {
		return nil
	}
	ans := []string{}
	for i := 1; i < n; i++ {
		if currentState[i] == '+' && currentState[i-1] == '+' {
			ans = append(ans, fmt.Sprintf("%s%s%s", currentState[:i-1], "--", currentState[i+1:]))
		}
	}
	return ans
}

// @lc code=end
