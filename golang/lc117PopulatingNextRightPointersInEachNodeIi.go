package golang

/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func lc117Connect(root *LinkedTreeNode) *LinkedTreeNode {
	start := root
	for start != nil {
		var newStart, prev *LinkedTreeNode
		handle := func(curr *LinkedTreeNode) {
			if curr == nil {
				return
			}
			if newStart == nil {
				newStart = curr
			}
			if prev != nil {
				prev.Next = curr
			}
			prev = curr
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = newStart
	}
	return root
}

// @lc code=end
