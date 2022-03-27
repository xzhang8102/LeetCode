package golang

import "fmt"

/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	n := len(nums)
	ans := []string{}
	for left, right := 0, 0; left < n; {
		if left == right || right < n && nums[right] == nums[right-1]+1 {
			right++
		} else {
			if right-left == 1 {
				ans = append(ans, fmt.Sprintf("%d", nums[left]))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", nums[left], nums[right-1]))
			}
			left = right
		}
	}
	return ans
}

// @lc code=end
