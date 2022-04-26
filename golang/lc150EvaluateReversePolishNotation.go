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
		switch size := len(stack); token {
		case "+", "-", "*", "/":
			op1, op2 := stack[size-2], stack[size-1]
			stack = stack[:size-2]
			res := op1 + op2
			if token == "-" {
				res = op1 - op2
			} else if token == "*" {
				res = op1 * op2
			} else if token == "/" {
				res = op1 / op2
			}
			stack = append(stack, res)
		default:
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		}
	}
	return stack[0]
}

// @lc code=end
