package golang

import "math"

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	pre := math.MinInt64
	stack := []interface{}{root}
	for len(stack) > 0 {
		ele := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node, ok := ele.(*TreeNode); ok {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			stack = append(stack, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			if val, _ := ele.(int); val <= pre {
				return false
			} else {
				pre = val
			}
		}
	}
	return true
}

// @lc code=end
