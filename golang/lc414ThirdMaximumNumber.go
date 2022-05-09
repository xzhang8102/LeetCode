package golang

/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] Third Maximum Number
 */

// @lc code=start
func thirdMax(nums []int) int {
	var first, second, third *int
	for _, num := range nums {
		n := num
		if first == nil || n > *first {
			first, second, third = &n, first, second
		} else if n < *first {
			if second == nil || n > *second {
				second, third = &n, second
			} else if n < *second {
				if third == nil || n > *third {
					third = &n
				}
			}
		}
	}
	if third == nil {
		return *first
	}
	return *third
}

// @lc code=end
