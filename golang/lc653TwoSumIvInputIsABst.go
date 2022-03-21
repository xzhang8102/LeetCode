package golang

/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
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
func findTarget(root *TreeNode, k int) bool {
	left, right := root, root
	leftStk := []*TreeNode{left}
	for left.Left != nil {
		leftStk = append(leftStk, left.Left)
		left = left.Left
	}
	rightStk := []*TreeNode{right}
	for right.Right != nil {
		rightStk = append(rightStk, right.Right)
		right = right.Right
	}
	for left != right {
		sum := left.Val + right.Val
		if sum == k {
			return true
		}
		if sum < k {
			left = leftStk[len(leftStk)-1]
			leftStk = leftStk[:len(leftStk)-1]
			for node := left.Right; node != nil; node = node.Left {
				leftStk = append(leftStk, node)
			}
		} else {
			right = rightStk[len(rightStk)-1]
			rightStk = rightStk[:len(rightStk)-1]
			for node := right.Left; node != nil; node = node.Right {
				rightStk = append(rightStk, node)
			}
		}
	}
	return false
}

// @lc code=end
