package golang

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + lc104Max(maxDepth(root.Left), maxDepth(root.Right))
}

func lc104Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
