package biweekly76

import "testing"

func Test_maximumScore(t *testing.T) {
	type args struct {
		scores []int
		edges  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case",
			args: args{
				scores: []int{14, 12, 10, 8, 1, 2, 3, 1},
				edges:  [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 5}, {1, 6}, {1, 7}, {1, 2}},
			},
			want: 44,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.args.scores, tt.args.edges); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
