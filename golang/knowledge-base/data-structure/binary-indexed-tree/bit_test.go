package main

import (
	"reflect"
	"testing"
)

func TestNewBIT(t *testing.T) {
	input := []int{3, 2, -1, 6, 5, 4, -3, 3, 7, 2, 3}
	want := &BIT{
		[]int{0, 3, 5, -1, 10, 5, 9, -3, 19, 7, 9, 3},
	}
	if got := NewBIT(input); !reflect.DeepEqual(got, want) {
		t.Errorf("NewBIT() = %v, want %v", got, want)
	}
	prefixSum := make([]int, len(input))
	prefixSum[0] = input[0]
	for i := 1; i < len(input); i++ {
		prefixSum[i] = prefixSum[i-1] + input[i]
	}
	for i := range prefixSum {
		if expected, sum := prefixSum[i], want.Query(i+1); sum != expected {
			t.Errorf("Prefix Sum from 0 to %d should be %d, but got %d", i, expected, sum)
		}
	}
}
