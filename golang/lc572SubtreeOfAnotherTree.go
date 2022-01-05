package golang

/*
 * @lc app=leetcode.cn id=572 lang=golang
 *
 * [572] Subtree of Another Tree
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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return lc572MatchTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func lc572MatchTree(root, target *TreeNode) bool {
	if root == nil && target == nil {
		return true
	}
	if root == nil || target == nil {
		return false
	}
	if root.Val != target.Val {
		return false
	}
	return lc572MatchTree(root.Left, target.Left) && lc572MatchTree(root.Right, target.Right)
}

// @lc code=end
