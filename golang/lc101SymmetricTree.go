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
	q := []*TreeNode{root, root}
	for len(q) > 0 {
		node1, node2 := q[0], q[1]
		q = q[2:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		q = append(q, node1.Right)
		q = append(q, node2.Left)
		q = append(q, node1.Left)
		q = append(q, node2.Right)
	}
	return true
}

// @lc code=end
