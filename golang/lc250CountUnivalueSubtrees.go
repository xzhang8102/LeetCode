package golang

/*
 * @lc app=leetcode.cn id=250 lang=golang
 *
 * [250] Count Univalue Subtrees
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
func countUnivalSubtrees(root *TreeNode) int {
	result := 0
	lc250Helper(root, &result)
	return result
}

func lc250Helper(root *TreeNode, result *int) {
	if root == nil {
		return
	}
	if isUniVal(root, root.Val) {
		*result++
	}
	lc250Helper(root.Left, result)
	lc250Helper(root.Right, result)
}

func isUniVal(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val == val {
		return isUniVal(root.Left, val) && isUniVal(root.Right, val)
	} else {
		return false
	}
}

// @lc code=end
