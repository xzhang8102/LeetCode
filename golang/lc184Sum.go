package golang

import "sort"

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	n := len(nums)
	for i := 0; i < n-3; i++ {
		if (i > 0 && nums[i] == nums[i-1]) || nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if (j > i+1 && nums[j] == nums[j-1]) || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					right--
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if sum < target {
					left++
					for left < right && nums[left] == nums[left-1] {
						left++
					}
				} else {
					right--
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				}
			}
		}
	}
	return ans
}

// @lc code=end
