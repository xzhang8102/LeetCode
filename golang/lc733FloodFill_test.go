package golang

import "testing"

func TestLC733(t *testing.T) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	floodFill(image, 1, 1, 2)
}
