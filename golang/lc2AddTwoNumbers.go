package golang

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ptr1, ptr2 := l1, l2
	ans := &ListNode{-1, nil}
	ansp := ans
	carry := 0
	for ptr1 != nil || ptr2 != nil || carry > 0 {
		val := 0
		if ptr1 != nil {
			val += ptr1.Val
			ptr1 = ptr1.Next
		}
		if ptr2 != nil {
			val += ptr2.Val
			ptr2 = ptr2.Next
		}
		val += carry
		carry = val / 10
		ansp.Next = &ListNode{val % 10, nil}
		ansp = ansp.Next
	}
	return ans.Next
}

// @lc code=end
