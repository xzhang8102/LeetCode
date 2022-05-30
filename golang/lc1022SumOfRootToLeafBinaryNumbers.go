package golang

/*
 * @lc app=leetcode.cn id=1022 lang=golang
 *
 * [1022] Sum of Root To Leaf Binary Numbers
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
func sumRootToLeaf(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum = sum*2 + node.Val
		if node.Left == nil && node.Right == nil {
			ans += sum
			return
		}
		if node.Left != nil {
			dfs(node.Left, sum)
		}
		if node.Right != nil {
			dfs(node.Right, sum)
		}
	}
	dfs(root, 0)
	return ans
}

// @lc code=end
