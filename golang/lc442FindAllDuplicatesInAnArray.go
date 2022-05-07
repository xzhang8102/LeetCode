package golang

/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] Find All Duplicates in an Array
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	n := len(nums)
	ans := []int{}
	for i := range nums {
		idx := (nums[i] - 1) % n
		nums[idx] += n
		if nums[idx] > 2*n {
			ans = append(ans, idx+1)
		}
	}
	return ans
}

// @lc code=end
