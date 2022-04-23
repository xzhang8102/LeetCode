package individual

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getNumber(root *TreeNode, ops [][]int) int {
	if root == nil {
		return 0
	}
	list := []int{}
	traverse(root, &list)
	ans := 0
	for i := len(ops) - 1; i >= 0; i-- {
		op := ops[i]
		begin := sort.SearchInts(list, op[1])
		end := sort.SearchInts(list, op[2])
		tmp := append([]int(nil), list[:begin]...)
		if end < len(list) {
			if list[end] == op[2] {
				tmp = append(tmp, list[end+1:]...)
			} else {
				tmp = append(tmp, list[end:]...)
			}
		}
		if op[0] == 1 {
			ans += len(list) - len(tmp)
		}
		list = tmp
	}
	return ans
}

func traverse(node *TreeNode, list *[]int) {
	if node == nil {
		return
	}
	traverse(node.Left, list)
	*list = append(*list, node.Val)
	traverse(node.Right, list)
}
