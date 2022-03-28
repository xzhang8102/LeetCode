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
	ptr := head
	ans := true
	var check func(node *ListNode)
	check = func(node *ListNode) {
		if node != nil {
			check(node.Next)
			if ptr.Val != node.Val {
				ans = false
			} else {
				ptr = ptr.Next
			}
		}
	}
	check(head)
	return ans
}

// @lc code=end
