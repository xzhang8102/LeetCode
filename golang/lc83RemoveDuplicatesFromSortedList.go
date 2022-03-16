package golang

/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	for head != nil {
		ptr := head.Next
		for ; ptr != nil && ptr.Val == head.Val; ptr = ptr.Next {
		}
		head.Next = ptr
		head = head.Next
	}
	return dummy.Next
}

// @lc code=end
