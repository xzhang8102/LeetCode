package golang

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, head}
	traverseList(dummy, n)
	return dummy.Next
}

func traverseList(head *ListNode, n int) int {
	if head == nil {
		return 0
	}
	count := traverseList(head.Next, n)
	if count == n {
		head.Next = head.Next.Next
	}
	return count + 1
}

// @lc code=end
