package golang

import (
	"fmt"
	"testing"
)

func TestIsCompleteTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				4,
				nil,
				nil,
			},
			Right: &TreeNode{
				5,
				nil,
				nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				6,
				nil,
				nil,
			},
			Right: nil,
		},
	}
	fmt.Println(isCompleteTree(root))
	fmt.Println(isCompleteTree1(root))
}
