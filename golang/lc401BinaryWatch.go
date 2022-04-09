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
	for h := 0; h <= 11; h++ {
		for m := 0; m <= 59; m++ {
			if bits.OnesCount(uint(h))+bits.OnesCount(uint(m)) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}

// @lc code=end
