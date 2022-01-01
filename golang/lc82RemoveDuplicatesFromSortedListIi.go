package golang

/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
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
	dummy := ListNode{
		Next: head,
	}
	ptr := &dummy
	for ptr != nil {
		p1, p2 := ptr.Next, ptr.Next
		for p2 != nil && p1.Val == p2.Val {
			p2 = p2.Next
		}
		if p1 != nil && p2 != p1.Next {
			ptr.Next = p2
		} else {
			ptr = ptr.Next
		}
	}
	return dummy.Next
}

// @lc code=end
