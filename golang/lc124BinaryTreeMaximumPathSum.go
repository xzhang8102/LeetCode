package golang

import "math"

/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt64
	var gain func(*TreeNode) int
	gain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := lc124Max(gain(node.Left), 0)
		right := lc124Max(gain(node.Right), 0)
		ans = lc124Max(ans, node.Val+left+right)
		return node.Val + lc124Max(left, right)
	}
	gain(root)
	return ans
}

func lc124Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
