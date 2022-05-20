package golang

/*
 * @lc app=leetcode.cn id=436 lang=golang
 *
 * [436] Find Right Interval
 */

// @lc code=start
func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	collect := map[int]int{}
	maxStart := int(-1e6 - 1)
	for i, interval := range intervals {
		start := interval[0]
		if start > maxStart {
			maxStart = start
		}
		collect[start] = i
	}
	for i := range ans {
		end := intervals[i][1]
		if end > maxStart {
			ans[i] = -1
		} else {
			for j := end; j <= maxStart; j++ {
				if start, ok := collect[j]; ok {
					ans[i] = start
					break
				}
			}
		}
	}
	return ans
}

// @lc code=end
