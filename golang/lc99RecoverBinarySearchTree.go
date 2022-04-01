package golang

import "math"

/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
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
func recoverTree(root *TreeNode) {
	prev := &TreeNode{Val: math.MinInt64}
	var err1, err2 *TreeNode

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev.Val >= node.Val && err1 == nil {
			err1 = prev
		}
		if prev.Val >= node.Val && err1 != nil {
			err2 = node
		}
		prev = node
		inorder(node.Right)
	}

	inorder(root)
	err1.Val, err2.Val = err2.Val, err1.Val
}

// @lc code=end
