package golang

/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > right {
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if left > interval[1] {
			ans = append(ans, interval)
		} else {
			left = lc57Min(left, interval[0])
			right = lc57Max(right, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return ans
}

func lc57Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lc57Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
