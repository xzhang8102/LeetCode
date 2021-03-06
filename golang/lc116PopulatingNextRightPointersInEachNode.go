package golang

/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] Populating Next Right Pointers in Each Node
 */

/**
 * Definition for a Node.
 * type Node struct {
 * 	 Val   int
 * 	 Left  *Node
 * 	 Right *Node
 * 	 Next  *Node
 * }
 */
type LinkedTreeNode struct {
	Val   int
	Left  *LinkedTreeNode
	Right *LinkedTreeNode
	Next  *LinkedTreeNode
}

// @lc code=start
func lc116Connect(root *LinkedTreeNode) *LinkedTreeNode {
	if root == nil {
		return nil
	}
	queue := []*LinkedTreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		tmp := queue
		queue = nil
		for i, node := range tmp {
			if i+1 < size {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

// @lc code=end
