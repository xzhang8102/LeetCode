package golang

/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N-ary Tree Level Order Traversal
 */

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type NaryTreeNode struct {
	Val      int
	Children []*NaryTreeNode
}

// @lc code=start
func levelOrder(root *NaryTreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	q := []*NaryTreeNode{root}
	for len(q) > 0 {
		level := []int{}
		tmp := q
		q = []*NaryTreeNode{}
		for _, node := range tmp {
			level = append(level, node.Val)
			q = append(q, node.Children...)
		}
		ans = append(ans, level)
	}
	return ans
}

// @lc code=end
