package crack

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}
	if root.Val == p.Val {
		if root.Right == nil {
			return nil
		}
		ptr := root.Right
		for ptr.Left != nil {
			ptr = ptr.Left
		}
		return ptr
	} else if root.Val > p.Val {
		next := inorderSuccessor(root.Left, p)
		if next == nil {
			return root
		} else {
			return next
		}
	} else {
		return inorderSuccessor(root.Right, p)
	}
}
