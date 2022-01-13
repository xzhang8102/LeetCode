package golang

/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] Arithmetic Slices
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	ans := 0
	gap, count := nums[0]-nums[1], 0
	// 因为等差数列的长度至少为 3，所以可以从 i=2 开始枚举
	/*
		input: [1,2,3,4,5]
								↑ [1,2,3] +1
									↑ [2,3,4], [1,2,3,4] +2
										↑ [2,3,4,5],[3,4,5],[1,2,3,4,5] +3
	*/
	for i := 2; i < n; i++ {
		if nums[i-1]-nums[i] == gap {
			count++
		} else {
			gap, count = nums[i-1]-nums[i], 0
		}
		ans += count
	}
	return ans
}

// @lc code=end
