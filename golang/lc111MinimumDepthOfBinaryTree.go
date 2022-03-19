package golang

import "math"

/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := math.MaxInt32
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth >= ans {
			return
		}
		if node.Left == nil && node.Right == nil {
			ans = lc111Min(ans, depth+1)
			return
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	return ans
}

func lc111Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
