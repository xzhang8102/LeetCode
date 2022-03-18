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
	ans := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		qSize := len(q)
		for qSize > 0 {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			qSize--
		}
		ans++
	}
	return ans
}

func lc104Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
