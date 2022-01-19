package golang

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	seen := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if index, ok := seen[nums[i]]; ok && i-index <= k {
			return true
		}
		seen[nums[i]] = i
	}
	return false
}

// @lc code=end
