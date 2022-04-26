package golang

import "strconv"

/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			op1, op2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			res := op1 + op2
			switch token {
			case "-":
				res = op1 - op2
			case "*":
				res = op1 * op2
			case "/":
				res = op1 / op2
			}
			stack = append(stack, res)
		}
	}
	return stack[0]
}

// @lc code=end
