package main

import (
	"reflect"
	"testing"
)

func Test_bubbleGrid_popBubbles(t *testing.T) {
	tests := []struct {
		name  string
		grid  bubbleGrid
		darts [][]int
		want  []int
	}{
		{
			name: "test case 1",
			grid: bubbleGrid{
				{1, 1, 0},
				{1, 0, 0},
				{1, 1, 0},
				{1, 1, 1},
			},
			darts: [][]int{
				{2, 0},
				{0, 0},
			},
			want: []int{4, 1},
		},
		{
			name: "test case 2",
			grid: bubbleGrid{
				{1, 0, 0, 0},
				{1, 1, 1, 0},
			},
			darts: [][]int{
				{1, 0},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.grid.popBubbles(tt.darts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleGrid.popBubbles() = %v, want %v", got, tt.want)
			}
		})
	}
}
