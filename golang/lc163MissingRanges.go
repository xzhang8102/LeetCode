package golang

import "fmt"

/*
 * @lc app=leetcode.cn id=163 lang=golang
 *
 * [163] Missing Ranges
 */

// @lc code=start
func findMissingRanges(nums []int, lower int, upper int) []string {
	ans := []string{}
	n := len(nums)
	for i := 0; i <= n; i++ {
		prev, curr := lower-1, upper+1
		if i > 0 {
			prev = nums[i-1]
		}
		if i < n {
			curr = nums[i]
		}
		if curr-prev == 2 {
			ans = append(ans, fmt.Sprintf("%d", curr-1))
		} else if curr-prev > 2 {
			ans = append(ans, fmt.Sprintf("%d->%d", prev+1, curr-1))
		}
	}
	return ans
}

// @lc code=end
