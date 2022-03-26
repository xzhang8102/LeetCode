package golang

/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}
	n := len(intervals)
	if n == 0 || newInterval[1] < intervals[0][0] {
		ans = append(ans, newInterval)
		ans = append(ans, intervals...)
		return ans
	}
	left, right := 0, 0
	for right < n {
		if intervals[left][1] < newInterval[0] {
			ans = append(ans, intervals[left])
			left++
			right++
		} else {
			if left == right || newInterval[1] >= intervals[right][0] || newInterval[1] >= intervals[right][1] {
				right++
			} else {
				break
			}
		}
	}
	lo := newInterval[0]
	if left < n {
		lo = lc57Min(newInterval[0], intervals[left][0])
	}
	hi := newInterval[1]
	merged := false
	if newInterval[1] >= intervals[right-1][0] {
		merged = true
		hi = lc57Max(newInterval[1], intervals[right-1][1])
	}
	ans = append(ans, []int{lo, hi})
	if merged {
		ans = append(ans, intervals[right:]...)
	} else {
		ans = append(ans, intervals[right-1:]...)
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
