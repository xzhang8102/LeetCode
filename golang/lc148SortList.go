package golang

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	h1 := sortList(slow.Next)
	slow.Next = nil
	h2 := sortList(head)
	dummy := &ListNode{Val: -1}
	ptr, p1, p2 := dummy, h1, h2
	for ; p1 != nil && p2 != nil; ptr = ptr.Next {
		if p1.Val < p2.Val {
			ptr.Next = p1
			p1 = p1.Next
		} else {
			ptr.Next = p2
			p2 = p2.Next
		}
	}
	if p1 != nil {
		ptr.Next = p1
	}
	if p2 != nil {
		ptr.Next = p2
	}
	return dummy.Next
}

// @lc code=end
