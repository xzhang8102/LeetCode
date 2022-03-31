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
	return lc98dfs(root.Left, math.MinInt64, root.Val) && lc98dfs(root.Right, root.Val, math.MaxInt64)
}

func lc98dfs(node *TreeNode, lo, hi int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lo || node.Val >= hi {
		return false
	}
	return lc98dfs(node.Left, lo, node.Val) && lc98dfs(node.Right, node.Val, hi)
}

// @lc code=end
