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
	ptrA, ptrB := headA, headB
	for ptrA != ptrB {
		if ptrA == nil {
			ptrA = headB
		} else {
			ptrA = ptrA.Next
		}
		if ptrB == nil {
			ptrB = headA
		} else {
			ptrB = ptrB.Next
		}
	}
	return ptrA
}

// @lc code=end
