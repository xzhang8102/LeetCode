package golang

import (
	"fmt"
	"math/bits"
)

/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] Binary Watch
 */

// @lc code=start
func readBinaryWatch(turnedOn int) []string {
	ans := []string{}
	for i := 0; i < 1024; i++ {
		h := i >> 6
		m := i & 0x03f
		if h < 12 && m < 60 && bits.OnesCount(uint(i)) == turnedOn {
			ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
		}
	}
	return ans
}

// @lc code=end
