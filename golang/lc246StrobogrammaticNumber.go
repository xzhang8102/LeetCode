package golang

/*
 * @lc app=leetcode.cn id=246 lang=golang
 *
 * [246] Strobogrammatic Number
 */

// @lc code=start
func isStrobogrammatic(num string) bool {
	n := len(num)
	rotated := make([]byte, n)
	for i, ch := range num {
		switch ch {
		case '1', '8', '0':
			rotated[n-1-i] = byte(ch)
		case '6':
			rotated[n-1-i] = '9'
		case '9':
			rotated[n-1-i] = '6'
		default:
			return false
		}
	}
	return string(rotated) == num
}

// @lc code=end
