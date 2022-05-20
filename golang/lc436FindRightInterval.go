package golang

import "sort"

/*
 * @lc app=leetcode.cn id=436 lang=golang
 *
 * [436] Find Right Interval
 */

// @lc code=start
func findRightInterval(intervals [][]int) []int {
	for i := range intervals {
		intervals[i] = append(intervals[i], i)
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	n := len(intervals)
	ans := make([]int, n)
	for _, interval := range intervals {
		idx := sort.Search(n, func(i int) bool { return intervals[i][0] >= interval[1] })
		if idx < n {
			ans[interval[2]] = intervals[idx][2]
		} else {
			ans[interval[2]] = -1
		}
	}
	return ans
}

// @lc code=end
