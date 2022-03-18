package golang

/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return lc108Helper(nums, 0, len(nums)-1)
}

func lc108Helper(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}
	mid := lo + (hi-lo)>>1
	root := &TreeNode{Val: nums[mid]}
	root.Left = lc108Helper(nums, lo, mid-1)
	root.Right = lc108Helper(nums, mid+1, hi)
	return root
}

// @lc code=end
