package golang

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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
func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	stack := []interface{}{}
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if val, ok := top.(int); ok {
			ans = append(ans, val)
		} else {
			node := top.(*TreeNode)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			stack = append(stack, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}
	return ans
}

// @lc code=end
