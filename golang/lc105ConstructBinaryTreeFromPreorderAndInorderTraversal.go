package golang

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	mapping := map[int]int{}
	for i, v := range inorder {
		mapping[v] = i
	}
	var dfs func(start, end int) *TreeNode
	dfs = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		val := preorder[0]
		preorder = preorder[1:]
		node := &TreeNode{Val: val}
		index := mapping[val]
		node.Left = dfs(start, index-1)
		node.Right = dfs(index+1, end)
		return node
	}
	return dfs(0, len(inorder)-1)
}

// @lc code=end
