package golang

import "testing"

func TestLC230(t *testing.T) {
	root := &TreeNode{
		3,
		&TreeNode{
			1,
			nil,
			&TreeNode{
				2,
				nil,
				nil,
			},
		},
		&TreeNode{
			4,
			nil,
			nil,
		},
	}
	kthSmallest(root, 1)
}
