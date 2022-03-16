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
	for root != nil {
		if root.Left == nil {
			ans = append(ans, root.Val)
			root = root.Right
		} else {
			preprocessor := root.Left
			for preprocessor.Right != nil && preprocessor.Right != root {
				preprocessor = preprocessor.Right
			}
			if preprocessor.Right == nil {
				preprocessor.Right = root
				root = root.Left
			} else {
				ans = append(ans, root.Val)
				preprocessor.Right = nil
				root = root.Right
			}
		}
	}
	return ans
}

// @lc code=end
