package golang

import "sort"

/*
 * @lc app=leetcode.cn id=954 lang=golang
 *
 * [954] Array of Doubled Pairs
 */

// @lc code=start
func canReorderDoubled(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool { return lc954Abs(arr[i]) <= lc954Abs(arr[j]) })
	cache := map[int]int{}
	for _, v := range arr {
		if v&1 == 0 && cache[v/2] > 0 {
			cache[v/2]--
		} else {
			cache[v]++
		}
	}
	for _, v := range cache {
		if v > 0 {
			return false
		}
	}
	return true
}

func lc954Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
