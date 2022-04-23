package team

import "testing"

func Test_conveyorBelt(t *testing.T) {
	type args struct {
		matrix []string
		start  []int
		end    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				matrix: []string{">>v", "v^<", "<><"}, start: []int{0, 1}, end: []int{2, 0},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conveyorBelt(tt.args.matrix, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("conveyorBelt() = %v, want %v", got, tt.want)
			}
		})
	}
}
