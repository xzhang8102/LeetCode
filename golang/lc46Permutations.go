package golang

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	n := len(nums)
	var result [][]int
	lc46Helper(0, n, nums, &result)
	return result
}

// count - the 'count'th slot we are filling
func lc46Helper(count, n int, nums []int, result *[][]int) {
	if count == n {
		tmp := append([]int{}, nums...)
		*result = append(*result, tmp)
		return
	}
	for i := count; i < n; i++ {
		nums[i], nums[count] = nums[count], nums[i]
		lc46Helper(count+1, n, nums, result)
		nums[i], nums[count] = nums[count], nums[i]
	}
}

// @lc code=end
