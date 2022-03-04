package golang

/*
 * @lc app=leetcode.cn id=2104 lang=golang
 *
 * [2104] Sum of Subarray Ranges
 */

// @lc code=start
func subArrayRanges(nums []int) int64 {
	n := len(nums)
	var ans int64 = 0
	for i := 0; i < n; i++ {
		min, max := nums[i], nums[i]
		for j := i + 1; j < n; j++ {
			if nums[j] > max {
				max = nums[j]
			}
			if nums[j] < min {
				min = nums[j]
			}
			ans += int64(max - min)
		}
	}
	return ans
}

// @lc code=end
