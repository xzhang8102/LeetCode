package golang

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	visited := map[*ListNode]bool{}
	for ptr := headA; ptr != nil; ptr = ptr.Next {
		visited[ptr] = true
	}
	for ptr := headB; ptr != nil; ptr = ptr.Next {
		if visited[ptr] {
			return ptr
		}
	}
	return nil
}

// @lc code=end
