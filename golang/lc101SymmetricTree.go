package golang

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return lc101IsMirror(root.Left, root.Right)
}

func lc101IsMirror(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return lc101IsMirror(p.Right, q.Left) && lc101IsMirror(p.Left, q.Right)
}

// @lc code=end
