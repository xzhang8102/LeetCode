package golang

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func lc102levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	level := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = []*TreeNode{}
		level++
		for _, node := range tmp {
			if node == nil {
				continue
			}
			if level > len(ans) {
				ans = append(ans, []int{})
			}
			ans[level-1] = append(ans[level-1], node.Val)
			q = append(q, node.Left)
			q = append(q, node.Right)
		}
	}
	return ans
}

// @lc code=end
