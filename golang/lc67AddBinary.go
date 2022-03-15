package golang

/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	n1, n2 := len(a), len(b)
	ans := []byte{}
	var carry byte
	i1, i2 := n1-1, n2-1
	for ; i1 >= 0 && i2 >= 0; i1, i2 = i1-1, i2-1 {
		d1, d2 := a[i1]-'0', b[i2]-'0'
		ans = append(ans, (d1+d2+carry)%2+'0')
		carry = (d1 + d2 + carry) / 2
	}
	for ; i1 >= 0; i1-- {
		d := a[i1] - '0'
		ans = append(ans, (d+carry)%2+'0')
		carry = (d + carry) / 2
	}
	for ; i2 >= 0; i2-- {
		d := b[i2] - '0'
		ans = append(ans, (d+carry)%2+'0')
		carry = (d + carry) / 2
	}
	if carry == 1 {
		ans = append(ans, '1')
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return string(ans)
}

// @lc code=end
