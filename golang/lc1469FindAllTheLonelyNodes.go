package golang

/*
 * @lc app=leetcode.cn id=1469 lang=golang
 *
 * [1469] Find All The Lonely Nodes
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
func getLonelyNodes(root *TreeNode) []int {
	ans := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right != nil {
			ans = append(ans, node.Right.Val)
		}
		if node.Right == nil && node.Left != nil {
			ans = append(ans, node.Left.Val)
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

// @lc code=end
