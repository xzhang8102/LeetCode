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
  result := make([][]int, 0)
  result = append(result, make([]int, 0))
  for i := range nums {
    backtrack(&result, nums, &[]int{nums[i]}, i)
  }
  return result
}

func backtrack(result *[][]int, nums []int, current *[]int ,index int) {
  *result = append(*result, *current)
  for i := index + 1; i < len(nums); i++ {
    tmp := make([]int, len(*current))
    copy(tmp, *current)
    tmp = append(tmp, nums[i])
    backtrack(result, nums, &tmp, i)
  }
}
// @lc code=end
