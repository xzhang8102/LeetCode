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
	val := root.Val
	q := []*TreeNode{root}
	for len(q) > 0 {
		head := q[0]
		q = q[1:]
		if head.Val != val {
			return false
		}
		if head.Left != nil {
			q = append(q, head.Left)
		}
		if head.Right != nil {
			q = append(q, head.Right)
		}
	}
	return true
}

// @lc code=end
