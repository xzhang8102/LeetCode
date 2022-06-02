package golang

/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] Delete Node in a BST
 */

// @lc code=start
func deleteNode(root *TreeNode, key int) *TreeNode {
	var curr, parent *TreeNode = root, nil
	for curr != nil && curr.Val != key {
		parent = curr
		if curr.Val > key {
			curr = curr.Left
		} else if curr.Val < key {
			curr = curr.Right
		}
	}
	if curr == nil {
		return root
	}
	if curr.Left == nil && curr.Right == nil {
		curr = nil
	} else if curr.Right == nil {
		curr = curr.Left
	} else if curr.Left == nil {
		curr = curr.Right
	} else {
		left := curr.Left
		curr = curr.Right
		ptr := curr
		for ptr.Left != nil {
			ptr = ptr.Left
		}
		ptr.Left = left
	}
	if parent == nil {
		return curr
	}
	if parent.Left != nil && parent.Left.Val == key {
		parent.Left = curr
	}
	if parent.Right != nil && parent.Right.Val == key {
		parent.Right = curr
	}
	return root
}

// @lc code=end
