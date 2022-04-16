package lccup2022spring

import "testing"

func Test_getNumber(t *testing.T) {
	type args struct {
		root *TreeNode
		ops  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				root: &TreeNode{
					4,
					&TreeNode{2, &TreeNode{Val: 1}, nil},
					&TreeNode{7, &TreeNode{5, nil, &TreeNode{Val: 6}}, nil},
				},
				ops: [][]int{{0, 2, 2}, {1, 1, 5}, {0, 4, 5}, {1, 5, 7}},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumber(tt.args.root, tt.args.ops); got != tt.want {
				t.Errorf("getNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
