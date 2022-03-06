package golang

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		dummy := &ListNode{}
		head := dummy
		ptr1, ptr2 := lists[0], lists[1]
		for ptr1 != nil && ptr2 != nil {
			if ptr1.Val <= ptr2.Val {
				head.Next = ptr1
				ptr1 = ptr1.Next
			} else {
				head.Next = ptr2
				ptr2 = ptr2.Next
			}
			head = head.Next
		}
		if ptr1 != nil {
			head.Next = ptr1
		}
		if ptr2 != nil {
			head.Next = ptr2
		}
		return dummy.Next
	}
	mid := len(lists) >> 1
	return mergeKLists([]*ListNode{mergeKLists(lists[:mid]), mergeKLists(lists[mid:])})
}

// @lc code=end
