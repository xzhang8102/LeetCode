package golang

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	ans := []string{}
	if n <= 0 {
		return ans
	}
	pick := []byte{}
	var backtrack func(int, int)
	backtrack = func(left, right int) {
		if left == 0 && right == 0 {
			ans = append(ans, string(append([]byte(nil), pick...)))
			return
		}
		if left > 0 {
			pick = append(pick, '(')
			backtrack(left-1, right)
			pick = pick[:len(pick)-1]
		}
		if right > 0 && left < right {
			pick = append(pick, ')')
			backtrack(left, right-1)
			pick = pick[:len(pick)-1]
		}
	}
	backtrack(n, n)
	return ans
}

// @lc code=end
