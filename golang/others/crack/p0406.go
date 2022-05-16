package crack

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if p == nil {
		return nil
	}
	var ans *TreeNode
	ptr := root
	for ptr != nil {
		if ptr.Val > p.Val {
			ans = ptr
			ptr = ptr.Left
		} else {
			ptr = ptr.Right
		}
	}
	return ans
}
