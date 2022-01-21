package golang

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=1182 lang=golang
 *
 * [1182] Shortest Distance to Target Color
 */

// @lc code=start
func shortestDistanceColor(colors []int, queries [][]int) []int {
	mapping := map[int][]int{}
	for index, color := range colors {
		mapping[color] = append(mapping[color], index)
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		colorArr := mapping[query[1]]
		if len(colorArr) == 0 {
			ans[i] = -1
			continue
		}
		idx := sort.SearchInts(colorArr, query[0])
		min := math.MaxInt32
		if idx > 0 {
			min = lc1182Min(min, query[0]-colorArr[idx-1])
		}
		if idx < len(colorArr) {
			min = lc1182Min(min, colorArr[idx]-query[0])
		}
		ans[i] = min
	}
	return ans
}

func lc1182Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
