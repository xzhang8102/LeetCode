package golang

import "container/list"

/*
 * @lc app=leetcode.cn id=2104 lang=golang
 *
 * [2104] Sum of Subarray Ranges
 */

// @lc code=start
func subArrayRanges(nums []int) int64 {
	n := len(nums)
	minRight := make([]int, n)
	stack := list.New()
	for i := 0; i < n; i++ {
		for top := stack.Back(); stack.Len() > 0 && nums[i] < nums[top.Value.(int)]; top = stack.Back() {
			minRight[top.Value.(int)] = i - top.Value.(int)
			stack.Remove(top)
		}
		stack.PushBack(i)
	}
	lc2104ClearStack(stack, minRight)
	maxRight := make([]int, n)
	for i := 0; i < n; i++ {
		for top := stack.Back(); stack.Len() > 0 && nums[i] > nums[top.Value.(int)]; top = stack.Back() {
			maxRight[top.Value.(int)] = i - top.Value.(int)
			stack.Remove(top)
		}
		stack.PushBack(i)
	}
	lc2104ClearStack(stack, maxRight)
	minLeft := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for top := stack.Back(); stack.Len() > 0 && nums[i] <= nums[top.Value.(int)]; top = stack.Back() {
			minLeft[top.Value.(int)] = top.Value.(int) - i
			stack.Remove(top)
		}
		stack.PushBack(i)
	}
	lc2104ClearStack(stack, minLeft)
	maxLeft := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for top := stack.Back(); stack.Len() > 0 && nums[i] >= nums[top.Value.(int)]; top = stack.Back() {
			maxLeft[top.Value.(int)] = top.Value.(int) - i
			stack.Remove(top)
		}
		stack.PushBack(i)
	}
	lc2104ClearStack(stack, maxLeft)
	var minSum, maxSum int64
	for i := 0; i < n; i++ {
		minSum += int64(minLeft[i] * minRight[i] * nums[i])
		maxSum += int64(maxLeft[i] * maxRight[i] * nums[i])
	}
	return maxSum - minSum
}

func lc2104ClearStack(stack *list.List, storage []int) {
	if stack.Len() > 0 {
		index := stack.Back().Value.(int)
		for stack.Len() > 0 {
			top := stack.Back()
			storage[top.Value.(int)] = lc2104Abs(index-top.Value.(int)) + 1
			stack.Remove(top)
		}
	}
}

func lc2104Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
