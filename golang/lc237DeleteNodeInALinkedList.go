package golang

/*
 * @lc app=leetcode.cn id=237 lang=golang
 *
 * [237] Delete Node in a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	curr, next := node, node.Next
	for next != nil {
		curr.Val = next.Val
		if next.Next == nil {
			curr.Next = nil
		}
		curr = next
		next = next.Next
	}
}

// @lc code=end
