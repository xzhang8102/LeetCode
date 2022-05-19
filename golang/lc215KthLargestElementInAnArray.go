package golang

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	for i := n / 2; i >= 0; i-- {
		lc215MaxHeapify(nums, i, n)
	}
	i := n - 1
	for i > n-k {
		nums[0], nums[i] = nums[i], nums[0]
		lc215MaxHeapify(nums, 0, i)
		i--
	}
	return nums[0]
}

func lc215MaxHeapify(nums []int, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && nums[child+1] > nums[child] {
			child++
		}
		if nums[root] >= nums[child] {
			return
		}
		nums[root], nums[child] = nums[child], nums[root]
		root = child
	}
}

// @lc code=end
