package golang

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] Subsets
 */

// solution with iteration
//
// func subsets(nums []int) [][]int {
// 	result := make([][]int, 0)
// 	result = append(result, make([]int, 0))
// 	for _, v := range nums {
// 		for i := range result {
// 			tmp := make([]int, len(result[i]))
// 			copy(tmp, result[i])
// 			tmp = append(tmp, v)
// 			result = append(result, tmp)
// 		}
// 	}
// 	return result
// }

// @lc code=start
func subsets(nums []int) [][]int {
	ans := [][]int{}
	var backtrack func(index int)
	pick := []int{}
	n := len(nums)
	backtrack = func(index int) {
		tmp := []int{}
		tmp = append(tmp, pick...)
		ans = append(ans, tmp)
		for i := index; i < n; i++ {
			pick = append(pick, nums[i])
			backtrack(i + 1)
			pick = pick[:len(pick)-1]
		}
	}
	backtrack(0)
	return ans
}

// @lc code=end
