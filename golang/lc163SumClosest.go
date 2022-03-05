package golang

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left, right := i+1, n-1; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			}
			if sum < target {
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else {
				right--
				for right > left && nums[right+1] == nums[right] {
					right--
				}
			}
			if gap := lc16Abs(target - sum); gap < lc16Abs(target-ans) {
				ans = sum
			}
		}
	}
	return ans
}

func lc16Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
