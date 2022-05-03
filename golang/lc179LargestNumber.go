package golang

import (
	"sort"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] Largest Number
 */

// @lc code=start
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		powx, powy := 10, 10
		for powx <= x {
			powx *= 10
		}
		for powy <= y {
			powy *= 10
		}
		return powy*x+y > powx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	var b strings.Builder
	for _, num := range nums {
		b.WriteString(strconv.Itoa(num))
	}
	return b.String()
}

// @lc code=end
