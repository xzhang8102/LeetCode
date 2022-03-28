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
	p1, p2 := lc235FindPath(root, p), lc235FindPath(root, q)
	for i := 0; i <= len(p1) && i <= len(p2); i++ {
		if i == len(p1) || i == len(p2) || p1[i] != p2[i] {
			return p1[i-1]
		}
	}
	return nil
}

func lc235FindPath(root, node *TreeNode) []*TreeNode {
	ret := []*TreeNode{root}
	for ptr := root; ptr != node; {
		if node.Val <= ptr.Val {
			ptr = ptr.Left
		} else {
			ptr = ptr.Right
		}
		ret = append(ret, ptr)
	}
	return ret
}

// @lc code=end
