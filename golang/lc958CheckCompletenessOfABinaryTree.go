package golang

/*
 * @lc app=leetcode.cn id=958 lang=golang
 *
 * [958] Check Completeness of a Binary Tree
 */

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 按照官方题解的思路
// 每一次层级的最后一个元素的编号跟当前层级之前所有元素的个数一样
func isCompleteTree1(root *TreeNode) bool {
	if root == nil {
		return false
	}
	type nTreeNode struct {
		n    int
		node *TreeNode
	}
	queue := make([]nTreeNode, 1)
	queue[0] = nTreeNode{1, root}
	size := 1
	for len(queue) > 0 {
		if size != queue[len(queue)-1].n {
			return false
		}
		tmp := queue
		queue = nil
		for _, item := range tmp {
			if item.node.Left != nil {
				size++
				queue = append(queue, nTreeNode{item.n * 2, item.node.Left})
			}
			if item.node.Right != nil {
				size++
				queue = append(queue, nTreeNode{item.n*2 + 1, item.node.Right})
			}
		}
	}
	return true
}

// @lc code=start
func isCompleteTree(root *TreeNode) bool {
	// found代表前一个元素是否是nil
	queue, found := []*TreeNode{root}, false
	for len(queue) > 0 {
		tmp := queue
		queue = nil
		for _, node := range tmp {
			if node == nil {
				found = true
			} else {
				if found {
					return false // 如果前一个元素为nil，后一个元素不为nil就违反了完全二叉树的规则
				}
				queue = append(queue, node.Left, node.Right)
			}
		}
	}
	return true
}

// @lc code=end
