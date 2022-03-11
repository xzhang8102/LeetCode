package golang

/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N-ary Tree Preorder Traversal
 */

type lc589Node struct {
	Val      int
	Children []*lc589Node
}

// @lc code=start
func preorder(root *lc589Node) []int {
	ans := []int{}
	stack := []*lc589Node{}
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return ans
}

// @lc code=end
