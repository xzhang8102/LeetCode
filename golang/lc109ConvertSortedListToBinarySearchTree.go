package golang

/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		if fast == nil || fast.Next == nil {
			tmp := slow.Next
			slow.Next = nil
			slow = tmp
		} else {
			slow = slow.Next
		}
	}
	root := &TreeNode{Val: slow.Val}
	if slow != head {
		root.Left = sortedListToBST(head)
	}
	root.Right = sortedListToBST(slow.Next)
	return root
}

// @lc code=end
