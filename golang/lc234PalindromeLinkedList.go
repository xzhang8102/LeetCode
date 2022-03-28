package golang

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	ans := true
	tail := lc234Reverse(slow.Next)
	for ptr1, ptr2 := head, tail; ptr1 != nil && ptr2 != nil; ptr1, ptr2 = ptr1.Next, ptr2.Next {
		if ptr1.Val != ptr2.Val {
			ans = false
		}
	}
	slow.Next = lc234Reverse(tail)
	return ans
}

func lc234Reverse(head *ListNode) *ListNode {
	var prev, curr *ListNode = nil, head
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}

// @lc code=end
