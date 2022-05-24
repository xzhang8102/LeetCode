package golang

/*
 * @lc app=leetcode.cn id=965 lang=golang
 *
 * [965] Univalued Binary Tree
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
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return (root.Left == nil || root.Left.Val == root.Val && isUnivalTree(root.Left)) && (root.Right == nil || root.Right.Val == root.Val && isUnivalTree(root.Right))
}

// @lc code=end
