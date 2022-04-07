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
	rootVal := postorder[len(postorder)-1]
	node := &TreeNode{Val: rootVal}
	inIdx := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			inIdx = i
		}
	}
	node.Left = buildTree(inorder[:inIdx], postorder[:inIdx])
	node.Right = buildTree(inorder[inIdx+1:], postorder[inIdx:len(postorder)-1])
	return node
}

// @lc code=end
