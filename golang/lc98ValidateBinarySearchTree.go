package golang

import "math"

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	pre := math.MinInt64
	var inorder func(node *TreeNode) bool
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !inorder(node.Left) {
			return false
		}
		if node.Val <= pre {
			return false
		}
		pre = node.Val
		if !inorder(node.Right) {
			return false
		}
		return true
	}
	return inorder(root)
}

// @lc code=end
