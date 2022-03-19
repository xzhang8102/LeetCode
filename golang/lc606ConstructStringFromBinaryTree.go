package golang

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=606 lang=golang
 *
 * [606] Construct String from Binary Tree
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
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var b strings.Builder
	stack := []interface{}{root}
	for len(stack) > 0 {
		ele := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node, ok := ele.(*TreeNode); ok {
			if node != root {
				b.WriteByte('(')
			}
			b.WriteString(strconv.Itoa(node.Val))
			if node.Left == nil && node.Right == nil {
				if node != root {
					b.WriteByte(')')
				}
				continue
			}
			if node != root {
				stack = append(stack, node.Val)
			}
			if node.Left == nil && node.Right != nil {
				b.WriteString("()")
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			b.WriteByte(')')
		}
	}
	return b.String()
}

// @lc code=end
