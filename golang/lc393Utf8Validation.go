package golang

/*
 * @lc app=leetcode.cn id=393 lang=golang
 *
 * [393] UTF-8 Validation
 */

// @lc code=start
func validUtf8(data []int) bool {
	n := len(data)
	for i := 0; i < n; {
		switch octet := data[i] & 0xff; {
		case octet>>7 == 0:
			i++
		case octet>>5 == 0x6:
			if i+1 < n {
				if (data[i+1]&0xff)>>6 == 0x2 {
					i += 2
				} else {
					return false
				}
			} else {
				return false
			}
		case octet>>4 == 0xe:
			if i+2 < n {
				for _, v := range data[i+1 : i+3] {
					if (v&0xff)>>6 != 0x2 {
						return false
					}
				}
				i += 3
			} else {
				return false
			}
		case octet>>3 == 0x1e:
			if i+3 < n {
				for _, v := range data[i+1 : i+4] {
					if (v&0xff)>>6 != 0x2 {
						return false
					}
				}
				i += 4
			} else {
				return false
			}
		default:
			return false
		}
	}
	return true
}

// @lc code=end
