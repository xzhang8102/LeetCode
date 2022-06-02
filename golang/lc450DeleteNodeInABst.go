package golang

/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] Delete Node in a BST
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			left := root.Left
			newNode := root.Right
			ptr := newNode
			for ptr.Left != nil {
				ptr = ptr.Left
			}
			ptr.Left = left
			return newNode
		}
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	} else {
		root.Left = deleteNode(root.Left, key)
		return root
	}
}

// @lc code=end
