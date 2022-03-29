package golang

import "sort"

/*
 * @lc app=leetcode.cn id=252 lang=golang
 *
 * [252] Meeting Rooms
 */

// @lc code=start
func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for j := 1; j < len(intervals); j++ {
		if intervals[j][0] < intervals[j-1][1] {
			return false
		}
	}
	return true
}

// @lc code=end
