package golang

/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] Path Sum
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	found := false
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil || found {
			return
		}
		if sum == node.Val && node.Left == nil && node.Right == nil {
			found = true
			return
		}
		dfs(node.Left, sum-node.Val)
		dfs(node.Right, sum-node.Val)
	}
	dfs(root, targetSum)
	return found
}

// @lc code=end
