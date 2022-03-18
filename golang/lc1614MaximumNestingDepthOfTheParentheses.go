package golang

/*
 * @lc app=leetcode.cn id=1614 lang=golang
 *
 * [1614] Maximum Nesting Depth of the Parentheses
 */

// @lc code=start
func lc1614maxDepth(s string) int {
	ans := 0
	n := len(s)
	ptr := 0 // index of left parenthesis
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			ptr++
			if ptr > ans {
				ans = ptr
			}
		}
		if s[i] == ')' {
			ptr--
		}
	}
	return ans
}

// @lc code=end
