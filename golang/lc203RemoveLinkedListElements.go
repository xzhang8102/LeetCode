package golang

/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{-1, head}
	for ptr := dummy; ptr != nil && ptr.Next != nil; {
		if ptr.Next.Val == val {
			next := ptr.Next
			ptr.Next = next.Next
			next.Next = nil
		} else {
			ptr = ptr.Next
		}
	}
	return dummy.Next
}

// @lc code=end
