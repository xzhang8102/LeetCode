package golang

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] Reorder List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	head2 := lc143Reverse(slow.Next)
	slow.Next = nil
	p, p1, p2 := head, head, head2
	for p1 != nil && p2 != nil {
		p1 = p1.Next
		p.Next = p2
		p = p.Next
		p2 = p2.Next
		p.Next = p1
		p = p.Next
	}
}

func lc143Reverse(node *ListNode) *ListNode {
	var prev, curr *ListNode = nil, node
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}
	return prev
}

// @lc code=end
