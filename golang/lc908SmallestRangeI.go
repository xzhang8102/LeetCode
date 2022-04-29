package golang

/*
 * @lc app=leetcode.cn id=908 lang=golang
 *
 * [908] Smallest Range I
 */

// @lc code=start
func smallestRangeI(nums []int, k int) int {
	min, max := nums[0], nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}
	if max-min <= 2*k {
		return 0
	} else {
		return max - min - 2*k
	}
}

// @lc code=end
