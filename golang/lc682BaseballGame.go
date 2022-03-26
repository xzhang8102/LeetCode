package golang

import "strconv"

/*
 * @lc app=leetcode.cn id=682 lang=golang
 *
 * [682] Baseball Game
 */

// @lc code=start
func calPoints(ops []string) int {
	n := len(ops)
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		switch ops[i] {
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		case "C":
			stack = stack[:len(stack)-1]
		case "D":
			stack = append(stack, stack[len(stack)-1]<<1)
		default:
			score, _ := strconv.Atoi(ops[i])
			stack = append(stack, score)
		}
	}
	ans := 0
	for _, v := range stack {
		ans += v
	}
	return ans
}

// @lc code=end
