package golang

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

/**
 * Definition for a Node.
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// @lc code=start
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	clone := &Node{}
	mapping := map[*Node]*Node{}
	for ptr1, ptr2 := head, clone; ptr1 != nil; ptr1, ptr2 = ptr1.Next, ptr2.Next {
		mapping[ptr1] = ptr2
		ptr2.Val = ptr1.Val
		if ptr1.Next != nil {
			ptr2.Next = &Node{}
		}
	}
	for ptr1, ptr2 := head, clone; ptr1 != nil; ptr1, ptr2 = ptr1.Next, ptr2.Next {
		if ptr1.Random != nil {
			ptr2.Random = mapping[ptr1.Random]
		}
	}
	return clone
}

// @lc code=end
