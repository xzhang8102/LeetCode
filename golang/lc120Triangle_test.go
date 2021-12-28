package golang

import "testing"

func TestLC120(t *testing.T) {
	input := [][]int{
		{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3},
	}
	minimumTotal(input)
}
