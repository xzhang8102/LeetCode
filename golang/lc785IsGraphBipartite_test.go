package golang

import (
	"fmt"
	"testing"
)

func TestLC785(t *testing.T) {
	graph := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	fmt.Println(isBipartite(graph))
}
