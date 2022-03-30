package golang

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	n := len(s)
	ans := []string{}
	pick := make([]int, 4, 4)
	var backtrack func(start, id int)
	backtrack = func(start, id int) {
		if id == 4 {
			if start == n {
				tmp := make([]string, 4)
				for i, v := range pick {
					tmp[i] = strconv.Itoa(v)
				}
				ans = append(ans, strings.Join(tmp, "."))
			}
			return
		}
		if start == n {
			return
		}
		if s[start] == '0' {
			pick[id] = 0
			backtrack(start+1, id+1)
		}
		addr := 0
		for i := start; i < n && i < start+3; i++ {
			addr = addr*10 + int(s[i]-'0')
			if addr > 0 && addr <= 0xFF {
				pick[id] = addr
				backtrack(i+1, id+1)
			} else {
				break
			}
		}
	}
	backtrack(0, 0)
	return ans
}

// @lc code=end
