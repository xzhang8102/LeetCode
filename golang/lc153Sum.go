package golang

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
import "sort"

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	sort.Ints(nums)
	for i, v := range nums {
		if i == 0 || v != nums[i-1] {
			for j, k := i+1, n-1; j < k; {
				if nums[j]+nums[k] == -v {
					ans = append(ans, []int{nums[i], nums[j], nums[k]})
					for j = j + 1; j < k && nums[j] == nums[j-1]; j++ {
					}
					for k = k - 1; j < k && nums[k] == nums[k+1]; k-- {
					}
				} else if nums[j]+nums[k] > -v {
					k--
				} else {
					j++
				}
			}
		}
	}
	return ans
}

// @lc code=end
