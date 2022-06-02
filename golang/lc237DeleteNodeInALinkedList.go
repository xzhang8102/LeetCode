package golang

/*
 * @lc app=leetcode.cn id=237 lang=golang
 *
 * [237] Delete Node in a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func lc237deleteNode(node *ListNode) {
	ptr := node.Next
	node.Val = ptr.Val
	node.Next = ptr.Next
	ptr.Next = nil
}

// @lc code=end
