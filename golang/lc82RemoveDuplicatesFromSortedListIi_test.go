package golang

import "testing"

func TestLC82(t *testing.T) {
	input := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					3,
					&ListNode{
						4,
						&ListNode{
							4,
							&ListNode{
								5, nil,
							},
						},
					},
				},
			},
		},
	}
	deleteDuplicates(input)
}
