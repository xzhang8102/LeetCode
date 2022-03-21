package golang

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	tortoise, hare := head, head.Next
	for hare != nil && hare.Next != nil {
		if hare == tortoise {
			return true
		}
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}
	return false
}

// @lc code=end
