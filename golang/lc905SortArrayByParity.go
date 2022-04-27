package golang

/*
 * @lc app=leetcode.cn id=905 lang=golang
 *
 * [905] Sort Array By Parity
 */

// @lc code=start
func sortArrayByParity(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i, j, k := 0, 0, n-1; i < n; i++ {
		if nums[i]&1 == 0 {
			ans[j] = nums[i]
			j++
		} else {
			ans[k] = nums[i]
			k--
		}
	}
	return ans
}

// @lc code=end
