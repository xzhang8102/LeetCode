package golang

import (
	"math"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=539 lang=golang
 *
 * [539] Minimum Time Difference
 */

// @lc code=start
func findMinDifference(timePoints []string) int {
	ans := math.MaxInt32
	n := len(timePoints)
	pick := []string{}
	var backtrack func(int)
	backtrack = func(index int) {
		if len(pick) == 2 {
			tmp := lc539Diff(pick[0], pick[1])
			if tmp < ans {
				ans = tmp
			}
			return
		}
		for i := index; i < n; i++ {
			pick = append(pick, timePoints[i])
			backtrack(i + 1)
			pick = pick[:len(pick)-1]
		}
	}
	backtrack(0)
	return ans
}

func lc539Diff(time1, time2 string) int {
	input1, input2 := lc539Parse(time1), lc539Parse(time2)
	if input1 > input2 {
		input1, input2 = input2, input1
	}
	diff1 := input2 - input1
	diff2 := input1 + lc539Parse("24:00") - input2
	if diff1 < diff2 {
		return diff1
	}
	return diff2
}

func lc539Parse(input string) int {
	minSec := strings.Split(input, ":")
	min, _ := strconv.Atoi(minSec[0])
	sec, _ := strconv.Atoi(minSec[1])
	return min*60 + sec
}

// @lc code=end
