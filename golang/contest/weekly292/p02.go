package weekly292

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode) (sum int, number int)
	dfs = func(node *TreeNode) (sum int, number int) {
		if node == nil {
			return
		}
		sumLeft, nLeft := dfs(node.Left)
		sumRight, nRight := dfs(node.Right)
		sum = node.Val + sumLeft + sumRight
		number = nLeft + nRight + 1
		if node.Val == sum/number {
			ans++
		}
		return
	}
	dfs(root)
	return ans
}
