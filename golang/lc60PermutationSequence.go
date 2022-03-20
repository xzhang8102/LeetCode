package golang

import "strings"

/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] Permutation Sequence
 */

// @lc code=start
func getPermutation(n int, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}
	k--
	pick := make([]int, n)
	var b strings.Builder
	for cnt := n - 1; cnt >= 0; cnt-- {
		offset := k/factorial[cnt] + 1
		k %= factorial[cnt]
		for i := 0; i < n; i++ {
			if pick[i] == 0 {
				offset--
				if offset == 0 {
					pick[i] = 1
					b.WriteByte(byte('1' + i))
					break
				}
			}
		}
	}
	return b.String()
}

// @lc code=end
