package golang

import "math"

/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] Third Maximum Number
 */

// @lc code=start
func thirdMax(nums []int) int {
	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > first {
			third = second
			second = first
			first = num
		} else if num != first {
			if num > second {
				third = second
				second = num
			} else if num != second {
				if num > third {
					third = num
				}
			}
		}
	}
	if third == math.MinInt64 {
		return first
	}
	return third
}

// @lc code=end
