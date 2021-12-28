package golang

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum1(nums []int, target int) []int {
	ans := make([]int, 2)
	cache := map[int]int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		wanted := target - nums[i]
		if val, seen := cache[wanted]; seen {
			ans[0], ans[1] = i, val
			return ans
		} else {
			cache[nums[i]] = i
		}
	}
	return ans
}

// @lc code=end
