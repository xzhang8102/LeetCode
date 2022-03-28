package golang

/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ptr := root
	for {
		if ptr.Val > p.Val && ptr.Val > q.Val {
			ptr = ptr.Left
		} else if ptr.Val < p.Val && ptr.Val < q.Val {
			ptr = ptr.Right
		} else {
			return ptr
		}
	}
}

// @lc code=end
