package golang

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] Reorder List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	nodes := []*ListNode{}
	for ptr := head; ptr != nil; ptr = ptr.Next {
		nodes = append(nodes, ptr)
	}
	n := len(nodes)
	for i, j := 0, n-1; i < n/2; i, j = i+1, j-1 {
		nodes[i].Next = nodes[j]
		if i+1 != j {
			nodes[j].Next = nodes[i+1]
			nodes[j-1].Next = nil
		} else {
			nodes[j].Next = nil
		}
	}
}

// @lc code=end
