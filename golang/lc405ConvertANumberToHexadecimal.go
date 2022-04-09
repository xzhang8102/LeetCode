package golang

/*
 * @lc app=leetcode.cn id=405 lang=golang
 *
 * [405] Convert a Number to Hexadecimal
 */

// @lc code=start
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	negetive := false
	if num < 0 {
		negetive = true
		num = -num
	}
	buf := [8]byte{}
	for i := 7; num > 0; i-- {
		buf[i] = byte(num % 16)
		num /= 16
	}
	if negetive {
		for i := range buf {
			buf[i] = ^buf[i]
			buf[i] &= 0x0f
		}
		for i := 7; i >= 0; i-- {
			buf[i]++
			if buf[i] < 16 {
				break
			} else {
				buf[i] = 0
			}
		}
	}
	i := 0
	for buf[i] == 0 {
		i++
	}
	ans := []byte{}
	for i < 8 {
		switch {
		case buf[i] >= 10:
			ans = append(ans, buf[i]-10+'a')
		default:
			ans = append(ans, buf[i]+'0')
		}
		i++
	}
	return string(ans)
}

// @lc code=end
