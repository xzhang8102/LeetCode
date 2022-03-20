package golang

/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
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
func findTarget(root *TreeNode, k int) bool {
	cache := map[int]bool{}
	found := false
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || found {
			return
		}
		if cache[k-node.Val] {
			found = true
			return
		}
		cache[node.Val] = true
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return found
}

// @lc code=end
