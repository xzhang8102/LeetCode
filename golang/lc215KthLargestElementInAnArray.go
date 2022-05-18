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
	lc215QuickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func lc215QuickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
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
	lc215QuickSort(nums, start, i)
	lc215QuickSort(nums, i+2, end)
}

// @lc code=end
