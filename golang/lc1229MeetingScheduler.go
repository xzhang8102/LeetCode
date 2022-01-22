package golang

import "sort"

/*
 * @lc app=leetcode.cn id=1229 lang=golang
 *
 * [1229] Meeting Scheduler
 */

// @lc code=start
func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	n1, n2 := len(slots1), len(slots2)
	sort.SliceStable(slots1, func(i, j int) bool {
		return slots1[i][0] < slots1[j][0]
	})
	sort.SliceStable(slots2, func(i, j int) bool {
		return slots2[i][0] < slots2[j][0]
	})
	for p1, p2 := 0, 0; p1 < n1 && p2 < n2; {
		lower := lc1229Max(slots1[p1][0], slots2[p2][0])
		overlap := lc1229Min(slots1[p1][1], slots2[p2][1]) - lower
		if overlap >= duration {
			return []int{lower, lower + duration}
		}
		if slots1[p1][1] >= slots2[p2][1] {
			p2++
		} else {
			p1++
		}
	}
	return []int{}
}

func lc1229Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lc1229Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
