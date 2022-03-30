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
	pick := make([]int, 0, 4)
	var backtrack func(start int)
	backtrack = func(start int) {
		if len(pick) == 4 && start == n {
			tmp := make([]string, 4)
			for i, idx := range pick {
				if i == 0 {
					tmp[i] = s[:idx]
				} else {
					tmp[i] = s[pick[i-1]:idx]
				}
			}
			ans = append(ans, strings.Join(tmp, "."))
		} else if len(pick) < 4 && n-start >= 4-len(pick) && n-start <= 3*(4-len(pick)) {
			for k := start; k < start+3 && k < n; k++ {
				switch k - start {
				case 0:
					pick = append(pick, k+1)
					backtrack(k + 1)
					pick = pick[:len(pick)-1]
				case 1:
					if s[k-1] != '0' {
						pick = append(pick, k+1)
						backtrack(k + 1)
						pick = pick[:len(pick)-1]
					}
				case 2:
					if val, _ := strconv.Atoi(s[start : k+1]); s[k-2] != '0' && val <= 255 {
						pick = append(pick, k+1)
						backtrack(k + 1)
						pick = pick[:len(pick)-1]
					}
				}
			}
		}
	}
	backtrack(0)
	return ans
}

// @lc code=end
