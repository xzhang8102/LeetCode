package golang

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
func swapPairs(head *ListNode) *ListNode {
	// find next head to start
	nextHead := head
	counter := 2
	for counter != 0 {
		if nextHead == nil {
			return head
		}
		counter--
		nextHead = nextHead.Next
	}
	// make sure the list is longer than 2
	var prev, next *ListNode
	curr := head
	for curr != nextHead && curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	// link head (new tail) to
	// swapPairs(nextHead)
	head.Next = swapPairs(nextHead)
	return prev
}

// @lc code=end
