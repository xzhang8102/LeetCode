package golang

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

// @lc code=start
func longestValidParentheses(s string) int {
	stack := []int{-1}
	ans := 0
	for i, char := range s {
		if char == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 1 && stack[0] == -1 {
				stack[0] = i
			} else if s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
				ans = lc32Max(ans, i-stack[len(stack)-1])
			} else {
				stack = append(stack, i)
			}
		}
	}
	return ans
}

func lc32Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
