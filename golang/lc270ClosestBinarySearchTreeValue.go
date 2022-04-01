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
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			continue
		}
		if diff := lc270Abs(float64(node.Val) - target); diff < min {
			min = diff
			ans = node.Val
		}
		if target > float64(node.Val) {
			q = append(q, node.Right)
		} else {
			q = append(q, node.Left)
		}
	}
	return ans
}

func lc270Abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
