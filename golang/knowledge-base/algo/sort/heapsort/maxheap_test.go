package heapsort

import (
	"math/rand"
	"sort"
	"testing"
)

type testData []int

func (t testData) Len() int           { return len(t) }
func (t testData) Less(i, j int) bool { return t[i] < t[j] }
func (t testData) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func Test_heapSort(t *testing.T) {
	n := 50
	round := 100
	input := make(testData, n)
	for r := 0; r < round; r++ {
		for i := range input {
			input[i] = rand.Intn(500)
		}
		heapSort(input, 0, 50)
		if !sort.IsSorted(input) {
			t.Error("Heap Sort Failed.")
		}
	}
}
