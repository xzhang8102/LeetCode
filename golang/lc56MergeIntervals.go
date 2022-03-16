package golang

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func lc56merge(intervals [][]int) [][]int {
	ans := [][]int{}
	n := len(intervals)
	if n == 0 {
		return ans
	}
	sort.Slice(intervals, func(i, j int) bool {
		loI, hiI := intervals[i][0], intervals[i][1]
		loJ, hiJ := intervals[j][0], intervals[j][1]
		if loI == loJ {
			return hiI < hiJ
		}
		return loI < loJ
	})
	for i := 0; i < n; i++ {
		lo, hi := intervals[i][0], intervals[i][1]
		if len(ans) == 0 || ans[len(ans)-1][1] < lo {
			ans = append(ans, []int{lo, hi})
		} else {
			ans[len(ans)-1][1] = lc56Max(ans[len(ans)-1][1], hi)
		}
	}
	return ans
}

func lc56Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
