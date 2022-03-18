package golang

/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		depthL := depth(root.Left)
		depthR := depth(root.Right)
		if depthL == -1 || depthR == -1 {
			return -1
		}
		if lc110Abs(depthL-depthR) > 1 {
			return -1
		}
		return 1 + lc110Max(depthL, depthR)
	}
	return depth(root) != -1
}

func lc110Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lc110Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
