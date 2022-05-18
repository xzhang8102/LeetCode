package golang

import (
	"math/rand"
	"time"
)

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return lc215QuickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func lc215QuickSelect(nums []int, start, end, target int) int {
	idx := lc215RandomPartition(nums, start, end)
	if idx == target {
		return nums[idx]
	} else if idx < target {
		return lc215QuickSelect(nums, idx+1, end, target)
	} else {
		return lc215QuickSelect(nums, start, idx-1, target)
	}
}

func lc215RandomPartition(nums []int, start, end int) int {
	idx := rand.Int()%(end-start+1) + start
	nums[idx], nums[end] = nums[end], nums[idx]
	val := nums[end]
	i := start - 1
	for j := start; j < end; j++ {
		if nums[j] <= val {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[end] = nums[end], nums[i+1]
	return i + 1
}

// @lc code=end
