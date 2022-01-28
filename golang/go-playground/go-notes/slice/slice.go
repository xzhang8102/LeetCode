package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	// fmt.Println(x[0]) panic: index out of range
	fmt.Println(len(x), cap(x))
	x = x[:5]
	fmt.Println(x[0], len(x), cap(x))
}
