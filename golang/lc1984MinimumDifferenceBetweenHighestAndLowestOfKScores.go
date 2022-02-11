package golang

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1984 lang=golang
 *
 * [1984] Minimum Difference Between Highest and Lowest of K Scores
 */

// @lc code=start
func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := nums[n-1] - nums[0]
	for i := k - 1; i < n; i++ {
		ans = lc1984Min(ans, nums[i]-nums[i-k+1])
	}
	return ans
}

func lc1984Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
