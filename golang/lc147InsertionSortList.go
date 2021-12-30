package golang

/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] Insertion Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	ptr := head.Next
	head.Next = nil
	for ptr != nil {
		p := dummy
		for p.Next != nil && ptr.Val > p.Next.Val {
			p = p.Next
		}
		nextPtr := ptr.Next
		ptr.Next = p.Next
		p.Next = ptr
		ptr = nextPtr
	}
	return dummy.Next
}

// @lc code=end
