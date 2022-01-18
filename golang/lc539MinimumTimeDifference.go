package golang

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=539 lang=golang
 *
 * [539] Minimum Time Difference
 */

// @lc code=start
func findMinDifference(timePoints []string) int {
	if len(timePoints) <= 1 || len(timePoints) > 1440 {
		return 0
	}
	sort.Strings(timePoints)
	ans := math.MaxInt32
	firstMins := lc539GetMinutes(timePoints[0])
	prevMins := firstMins
	for _, t := range timePoints[1:] {
		mins := lc539GetMinutes(t)
		ans = lc539Min(ans, mins-prevMins)
		prevMins = mins
	}
	ans = lc539Min(ans, firstMins+1440-prevMins)
	return ans
}

func lc539GetMinutes(input string) int {
	return (int(input[0]-'0')*10+int(input[1]-'0'))*60 + int(input[3]-'0')*10 + int(input[4]-'0')
}

func lc539Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
