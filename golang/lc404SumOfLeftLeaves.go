package golang

/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] Sum of Left Leaves
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
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		for _, node := range q[:size] {
			if node.Left != nil {
				if node.Left.Left == nil && node.Left.Right == nil {
					ans += node.Left.Val
				} else {
					q = append(q, node.Left)
				}
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[size:]
	}
	return ans
}

// @lc code=end
