package golang

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	return lc206Helper(head, nil)
}

func lc206Helper(curr, prev *ListNode) *ListNode {
	if curr == nil {
		return curr
	}
	newCurr := curr.Next
	curr.Next = prev
	if newCurr == nil {
		return curr
	}
	return lc206Helper(newCurr, curr)
}

// @lc code=end
