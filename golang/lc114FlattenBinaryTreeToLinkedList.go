package golang

/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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
func flatten(root *TreeNode) {
	ptr := root
	for ptr != nil {
		if ptr.Left != nil {
			rightmost := ptr.Left
			for rightmost.Right != nil {
				rightmost = rightmost.Right
			}
			rightmost.Right = ptr.Right
			ptr.Right = ptr.Left
			ptr.Left = nil
		}
		ptr = ptr.Right
	}
}

// @lc code=end
