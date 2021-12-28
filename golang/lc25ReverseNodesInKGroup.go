package golang

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func reverseKGroup(head *ListNode, k int) *ListNode {
  nextHead := head
  counter := k
  for counter != 0 {
    if nextHead == nil {
      return head
    }
    counter --
    nextHead = nextHead.Next
  }
  var prev, next *ListNode
  curr := head
  for curr != nextHead && curr != nil {
    next = curr.Next
    curr.Next = prev
    prev = curr
    curr = next
  }
  head.Next = reverseKGroup(nextHead, k)
  return prev
}

// @lc code=end
