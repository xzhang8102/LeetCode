package golang

import "testing"

func TestLC116(t *testing.T) {
	root := &Node{
		1,
		&Node{
			2,
			&Node{
				4,
				nil,
				nil,
				nil,
			},
			&Node{
				5,
				nil,
				nil,
				nil,
			},
			nil,
		},
		&Node{
			3,
			&Node{
				6,
				nil,
				nil,
				nil,
			},
			&Node{
				7,
				nil,
				nil,
				nil,
			},
			nil,
		},
		nil,
	}
	connect(root)
}
