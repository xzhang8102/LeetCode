package golang

import "sort"

/*
 * @lc app=leetcode.cn id=2055 lang=golang
 *
 * [2055] Plates Between Candles
 */

// @lc code=start
func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	// assume the left end is '|'
	sum := make([]int, n+1)
	candleIdx := []int{}
	ans := make([]int, len(queries))
	for i := 0; i < n; i++ {
		if s[i] == '*' {
			sum[i+1] = sum[i] + 1
		} else {
			candleIdx = append(candleIdx, i)
			sum[i+1] = sum[i]
		}
	}
	if len(candleIdx) == 0 {
		return ans
	}
	for idx, query := range queries {
		// find the closest '|'
		leftIdx := sort.SearchInts(candleIdx, query[0])
		if leftIdx == len(candleIdx) {
			ans[idx] = 0
			continue
		}
		left := candleIdx[leftIdx]
		rightIdx := sort.SearchInts(candleIdx, query[1])
		right := query[1]
		if rightIdx == len(candleIdx) {
			right = candleIdx[rightIdx-1]
		} else if candleIdx[rightIdx] != right {
			if rightIdx == 0 {
				ans[idx] = 0
				continue
			} else {
				right = candleIdx[rightIdx-1]
			}
		}
		if left <= right {
			ans[idx] = sum[right] - sum[left]
		} else {
			ans[idx] = 0
		}
	}
	return ans
}

// @lc code=end
