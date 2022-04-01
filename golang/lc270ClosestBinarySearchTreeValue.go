package golang

import "math"

/*
 * @lc app=leetcode.cn id=270 lang=golang
 *
 * [270] Closest Binary Search Tree Value
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
func closestValue(root *TreeNode, target float64) int {
	ans := 0
	min := math.MaxFloat64
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if min == 0 || node == nil {
			return
		}
		if diff := lc270Abs(float64(node.Val) - target); diff < min {
			min = diff
			ans = node.Val
		}
		if target < float64(node.Val) {
			dfs(node.Left)
		} else {
			dfs(node.Right)
		}
	}
	dfs(root)
	return ans
}

func lc270Abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func lc270Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
