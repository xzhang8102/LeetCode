package golang

import "math"

/*
 * @lc app=leetcode.cn id=1182 lang=golang
 *
 * [1182] Shortest Distance to Target Color
 */

// @lc code=start
func shortestDistanceColor(colors []int, queries [][]int) []int {
	n := len(colors)
	left, right := make([][3]int, n), make([][3]int, n)
	for i := range left {
		left[i] = [3]int{-1, -1, -1}
		right[i] = [3]int{-1, -1, -1}
	}
	for i := 0; i < n; i++ {
		for c := 0; c < 3; c++ {
			if i > 0 && left[i-1][c] != -1 {
				left[i][c] = left[i-1][c] + 1
			}
		}
		left[i][colors[i]-1] = 0
	}
	for i := n - 1; i >= 0; i-- {
		for c := 0; c < 3; c++ {
			if i < n-1 && right[i+1][c] != -1 {
				right[i][c] = right[i+1][c] + 1
			}
		}
		right[i][colors[i]-1] = 0
	}
	ans := []int{}
	for _, q := range queries {
		disLeft, disRight := left[q[0]][q[1]-1], right[q[0]][q[1]-1]
		min := math.MaxInt32
		if disLeft != -1 {
			min = lc1182Min(min, disLeft)
		}
		if disRight != -1 {
			min = lc1182Min(min, disRight)
		}
		if min == math.MaxInt32 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, min)
		}
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
