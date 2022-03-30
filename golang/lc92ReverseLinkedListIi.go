package golang

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left >= right || head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{-1, head}
	var prev, curr *ListNode = dummy, head
	var tail, newHead *ListNode
	for right > 1 {
		if left > 1 {
			prev = prev.Next
			curr = curr.Next
		} else {
			if left == 1 {
				tail = curr
				newHead = prev
			}
			tmp := curr
			curr = curr.Next
			tmp.Next = prev
			prev = tmp
		}
		right--
		left--
	}
	tail.Next = curr.Next
	curr.Next = prev
	newHead.Next = curr
	return dummy.Next
}

// @lc code=end
