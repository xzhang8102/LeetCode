package golang

/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] Base 7
 */

// @lc code=start
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	buffer := []byte{}
	negative := false
	if num < 0 {
		negative = true
		num = -num
	}
	for num > 0 {
		buffer = append(buffer, byte('0'+num%7))
		num /= 7
	}
	if negative {
		buffer = append(buffer, '-')
	}
	for i := 0; i < len(buffer)/2; i++ {
		buffer[i], buffer[len(buffer)-1-i] = buffer[len(buffer)-1-i], buffer[i]
	}
	return string(buffer)
}

// @lc code=end
