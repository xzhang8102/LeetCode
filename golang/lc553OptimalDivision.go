package golang

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=553 lang=golang
 *
 * [553] Optimal Division
 */

// @lc code=start
func optimalDivision(nums []int) string {
	var builder strings.Builder
	n := len(nums)
	if n == 1 {
		return fmt.Sprintf("%d", nums[0])
	}
	if n == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}
	for i, num := range nums {
		if i == 0 {
			builder.WriteString(fmt.Sprintf("%d/(", num))
		} else if i == n-1 {
			builder.WriteString(fmt.Sprintf("%d)", num))
		} else {
			builder.WriteString(fmt.Sprintf("%d/", num))
		}
	}
	return builder.String()
}

// @lc code=end
