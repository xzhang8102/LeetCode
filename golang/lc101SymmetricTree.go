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
	q1, q2 := []*TreeNode{}, []*TreeNode{}
	if root.Left != nil {
		q1 = append(q1, root.Left)
	}
	if root.Right != nil {
		q2 = append(q2, root.Right)
	}
	for len(q1) > 0 && len(q2) > 0 {
		node1, node2 := q1[0], q2[0]
		q1, q2 = q1[1:], q2[1:]
		if node1.Val != node2.Val {
			return false
		}
		if node1.Right == nil && node2.Left != nil {
			return false
		}
		if node1.Left == nil && node2.Right != nil {
			return false
		}
		if node1.Right != nil {
			q1 = append(q1, node1.Right)
		}
		if node2.Left != nil {
			q2 = append(q2, node2.Left)
		}
		if node1.Left != nil {
			q1 = append(q1, node1.Left)
		}
		if node2.Right != nil {
			q2 = append(q2, node2.Right)
		}
	}
	return len(q1) == 0 && len(q2) == 0
}

// @lc code=end
