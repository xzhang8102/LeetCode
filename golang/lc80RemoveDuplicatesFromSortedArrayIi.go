package golang

/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	n := len(nums)
	left := 0
	cnt := 0
	for right := 0; right <= n; {
		if left == right || right < n && nums[right] == nums[left] {
			cnt++
			right++
		} else {
			nums[left] = nums[right-1]
			if cnt > 1 {
				left++
				nums[left] = nums[right-1]
			}
			left++
			if right < n {
				nums[left] = nums[right]
			} else {
				break
			}
			cnt = 0
		}
	}
	return left
}

// @lc code=end
