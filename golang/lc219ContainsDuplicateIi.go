package golang

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	window := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		if i > k {
			delete(window, nums[i-k-1])
		}
		if _, ok := window[nums[i]]; ok {
			return true
		}
		window[nums[i]] = struct{}{}
	}
	return false
}

// @lc code=end
