package golang

/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head.Next, head.Next.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	ans := head
	for ans != fast {
		ans = ans.Next
		fast = fast.Next
	}
	return ans
}

// @lc code=end
