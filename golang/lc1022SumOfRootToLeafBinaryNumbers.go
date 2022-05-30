package golang

/*
 * @lc app=leetcode.cn id=1022 lang=golang
 *
 * [1022] Sum of Root To Leaf Binary Numbers
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
func sumRootToLeaf(root *TreeNode) (ans int) {
	if root == nil {
		return
	}
	q := []struct {
		node *TreeNode
		sum  int
	}{{root, root.Val}}
	for len(q) > 0 {
		head := q[0]
		q = q[1:]
		if head.node.Left == nil && head.node.Right == nil {
			ans += head.sum
		}
		if head.node.Left != nil {
			q = append(q, struct {
				node *TreeNode
				sum  int
			}{head.node.Left, head.sum*2 + head.node.Left.Val})
		}
		if head.node.Right != nil {
			q = append(q, struct {
				node *TreeNode
				sum  int
			}{head.node.Right, head.sum*2 + head.node.Right.Val})
		}
	}
	return ans
}

// @lc code=end
