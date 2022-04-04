package main

import (
	"fmt"
	"math/rand"
)

func main() {
	input := make([]int, 50)
	for i := range input {
		input[i] = rand.Intn(100)
	}
	// fmt.Println(MergeSortOut(input))
	fmt.Println(input)
	MergeSortIn(input, 0, len(input)-1)
	fmt.Println(input)
}

// inplace sort
func MergeSortIn(input []int, start int, end int) {
	if start >= end {
		return
	}
	mid := (start + end) >> 1
	MergeSortIn(input, start, mid)
	MergeSortIn(input, mid+1, end)
	// no need further sorting
	if input[mid] <= input[mid+1] {
		return
	}
	l, r := start, mid+1
	for l <= mid && r <= end {
		if input[l] <= input[r] {
			l++
		} else {
			tmp := input[r]          // take the smaller one out
			for k := r; k > l; k-- { // rotate elements between l and r by 1
				input[k] = input[k-1]
			}
			input[l] = tmp
			l++
			r++
			mid++
		}
	}
}

// outplace sort
func MergeSortOut(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	mid := len(input) >> 1
	left := MergeSortOut(input[:mid])
	right := MergeSortOut(input[mid:])
	buffer := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			buffer = append(buffer, left[i])
			i++
		} else {
			buffer = append(buffer, right[j])
			j++
		}
	}
	buffer = append(buffer, left[i:]...)
	buffer = append(buffer, right[j:]...)
	return buffer
}
