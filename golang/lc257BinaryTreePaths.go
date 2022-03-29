package golang

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] Binary Tree Paths
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
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	ans := []string{}
	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node.Left != nil {
			dfs(node.Left, fmt.Sprintf("%s->%d", path, node.Left.Val))
		}
		if node.Right != nil {
			dfs(node.Right, fmt.Sprintf("%s->%d", path, node.Right.Val))
		}
		if node.Left == nil && node.Right == nil {
			ans = append(ans, path)
		}
	}
	dfs(root, strconv.Itoa(root.Val))
	return ans
}

// @lc code=end
