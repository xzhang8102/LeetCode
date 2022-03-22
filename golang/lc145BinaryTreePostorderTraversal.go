package golang

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := []int{}
	stack := []interface{}{root}
	for len(stack) > 0 {
		ele := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node, ok := ele.(*TreeNode); ok {
			stack = append(stack, node.Val)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			ans = append(ans, ele.(int))
		}
	}
	return ans
}

// @lc code=end
