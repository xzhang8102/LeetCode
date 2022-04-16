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
				scores: []int{8, 7, 1, 26},
				edges:  [][]int{{3, 1}, {1, 2}, {2, 0}, {0, 3}, {2, 3}, {0, 1}},
			},
			want: 42,
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
