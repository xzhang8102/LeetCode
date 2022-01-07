package golang

/*
 * @lc app=leetcode.cn id=1614 lang=golang
 *
 * [1614] Maximum Nesting Depth of the Parentheses
 */

// @lc code=start
func maxDepth(s string) int {
	ans := 0
	n := len(s)
	if n == 0 {
		return ans
	}
	stack := []byte{}
	ptr := -1 // index of left parenthesis
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
			ptr++
			if ptr+1 > ans {
				ans = ptr + 1
			}
		}
		if s[i] == ')' {
			if len(stack) == 0 {
				return 0
			}
			stack = stack[:ptr]
			ptr--
		}
	}
	if len(stack) > 0 {
		return 0
	}
	return ans
}

// @lc code=end
