package golang

import "sort"

/*
 * @lc app=leetcode.cn id=668 lang=golang
 *
 * [668] Kth Smallest Number in Multiplication Table
 */

// @lc code=start
func findKthNumber(m int, n int, k int) int {
	return sort.Search(m*n+1, func(val int) bool {
		cnt := val / n * n
		for i := val/n + 1; i <= m; i++ {
			cnt += val / i
		}
		return cnt >= k
	})
}

// @lc code=end
