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
	hh := [5][]int{}
	mm := [7][]int{}
	for i := 0; i <= 59; i++ {
		n := bits.OnesCount(uint(i))
		if i <= 11 {
			hh[n] = append(hh[n], i)
		}
		mm[n] = append(mm[n], i)
	}
	for h := 0; h <= 4 && h <= turnedOn; h++ {
		m := turnedOn - h
		if m > 6 {
			continue
		}
		for _, hv := range hh[h] {
			for _, mv := range mm[m] {
				ans = append(ans, fmt.Sprintf("%d:%02d", hv, mv))
			}
		}
	}
	return ans
}

// @lc code=end
