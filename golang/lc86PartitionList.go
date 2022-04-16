package golang

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] Partition List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func lc86partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead, dummy := &ListNode{Val: -1}, &ListNode{Val: x}
	ptr, ptr1, ptr2 := head, dummyHead, dummy
	for ptr != nil {
		if ptr.Val < x {
			ptr1.Next = ptr
			ptr1 = ptr1.Next
		} else {
			ptr2.Next = ptr
			ptr2 = ptr2.Next
		}
		ptr = ptr.Next
	}
	ptr2.Next = nil
	ptr1.Next = dummy.Next
	dummy.Next = nil
	return dummyHead.Next
}

// @lc code=end
