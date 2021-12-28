package golang

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	cursor := -1
	stack := make([]rune, len(s))
	for _, c := range s {
		switch c {
		case ')', ']', '}':
			if cursor < 0 {
				return false
			}
			if !validHelper(c, stack[cursor]) {
				return false
			}
			cursor--
		default:
			cursor++
			stack[cursor] = c
		}
	}
	return cursor == -1
}

func validHelper(input rune, top rune) bool {
	switch input {
	case ']':
		return top == '['
	case '}':
		return top == '{'
	case ')':
		return top == '('
	default:
		return false
	}
}

// @lc code=end
