package golang

/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	stack := []*TreeNode{root}
	inIdx := len(inorder) - 1
	for i := len(postorder) - 2; i >= 0; i-- {
		postVal := postorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inIdx] {
			node.Right = &TreeNode{Val: postVal}
			stack = append(stack, node.Right)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inIdx] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inIdx--
			}
			node.Left = &TreeNode{Val: postVal}
			stack = append(stack, node.Left)
		}
	}
	return root
}

// @lc code=end
