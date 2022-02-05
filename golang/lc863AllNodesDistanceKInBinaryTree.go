package golang

/*
 * @lc app=leetcode.cn id=863 lang=golang
 *
 * [863] All Nodes Distance K in Binary Tree
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
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parents := map[int]*TreeNode{}
	var findParent func(*TreeNode)
	findParent = func(node *TreeNode) {
		if node.Left != nil {
			parents[node.Left.Val] = node
			findParent(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			findParent(node.Right)
		}
	}
	findParent(root)
	ans := []int{}
	var traverse func(node, from *TreeNode, depth int)
	traverse = func(node, from *TreeNode, depth int) {
		if depth == k {
			ans = append(ans, node.Val)
			return
		}
		if node.Left != nil && node.Left != from {
			traverse(node.Left, node, depth+1)
		}
		if node.Right != nil && node.Right != from {
			traverse(node.Right, node, depth+1)
		}
		if parent := parents[node.Val]; parent != nil && parent != from {
			traverse(parent, node, depth+1)
		}
	}
	traverse(target, nil, 0)
	return ans
}

// @lc code=end
