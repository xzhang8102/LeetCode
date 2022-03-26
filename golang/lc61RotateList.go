package golang

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	length := 1
	tail := head
	for tail.Next != nil {
		length++
		tail = tail.Next
	}
	k %= length
	if k == 0 {
		return head
	}
	ptr := head
	for cnt := length - k; cnt > 1; cnt, ptr = cnt-1, ptr.Next {
	}
	ans := ptr.Next
	ptr.Next = nil
	tail.Next = head
	return ans
}

// @lc code=end
