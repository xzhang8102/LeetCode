package golang

import "math"

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
func lc105buildTree(preorder []int, inorder []int) *TreeNode {
	preIdx, inIdx := 0, 0
	var dfs func(stop int) *TreeNode
	dfs = func(stop int) *TreeNode {
		if preIdx >= len(preorder) {
			return nil
		}
		if inorder[inIdx] == stop {
			inIdx++
			return nil
		}
		node := &TreeNode{Val: preorder[preIdx]}
		preIdx++
		node.Left = dfs(node.Val)
		node.Right = dfs(stop)
		return node
	}
	return dfs(math.MaxInt32)
}

// @lc code=end
