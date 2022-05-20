package golang

import "sort"

/*
 * @lc app=leetcode.cn id=436 lang=golang
 *
 * [436] Find Right Interval
 */

// @lc code=start
func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	sortStart, sortEnd := make([]struct{ val, i int }, n), make([]struct{ val, i int }, n)
	for i, interval := range intervals {
		sortStart[i] = struct {
			val int
			i   int
		}{interval[0], i}
		sortEnd[i] = struct {
			val int
			i   int
		}{interval[1], i}
	}
	sort.Slice(sortStart, func(i, j int) bool { return sortStart[i].val < sortStart[j].val })
	sort.Slice(sortEnd, func(i, j int) bool { return sortEnd[i].val < sortEnd[j].val })
	ans := make([]int, n)
	ptr := 0
	for _, p := range sortEnd {
		for ptr < n && sortStart[ptr].val < p.val {
			ptr++
		}
		if ptr < n {
			ans[p.i] = sortStart[ptr].i
		} else {
			ans[p.i] = -1
		}
	}
	return ans
}

// @lc code=end
