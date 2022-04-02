package golang

/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	dir := 1
	for len(q) > 0 {
		tmp := q
		q = []*TreeNode{}
		ans = append(ans, make([]int, len(tmp)))
		i := 0
		if dir == -1 {
			i = len(tmp) - 1
		}
		for _, node := range tmp {
			ans[len(ans)-1][i] = node.Val
			if dir > 0 {
				i++
			} else {
				i--
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		dir *= -1
	}
	return ans
}

// @lc code=end
