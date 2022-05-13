package golang

/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	record := map[int]*TreeNode{}
	depth := -1
	stack := []struct {
		node  *TreeNode
		depth int
	}{{root, 0}}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		node, d := top.node, top.depth
		stack = stack[:len(stack)-1]
		if d > depth {
			depth = d
		}
		if _, ok := record[d]; !ok {
			record[d] = node
		}
		if node.Left != nil {
			stack = append(stack, struct {
				node  *TreeNode
				depth int
			}{node.Left, d + 1})
		}
		if node.Right != nil {
			stack = append(stack, struct {
				node  *TreeNode
				depth int
			}{node.Right, d + 1})
		}
	}
	for i := 0; i <= depth; i++ {
		ans = append(ans, record[i].Val)
	}
	return ans
}

// @lc code=end
