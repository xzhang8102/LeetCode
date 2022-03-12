package golang

/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N-ary Tree Postorder Traversal
 */

type lc590Node struct {
	Val      int
	Children []*lc590Node
}

// @lc code=start

func postorder(root *lc590Node) []int {
	ans := []int{}
	var traverse func(root *lc590Node)
	traverse = func(root *lc590Node) {
		if root == nil {
			return
		}
		for _, child := range root.Children {
			traverse(child)
		}
		ans = append(ans, root.Val)
	}
	traverse(root)
	return ans
}

// @lc code=end
