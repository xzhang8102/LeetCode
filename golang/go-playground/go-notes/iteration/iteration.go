package main

import "fmt"

func main() {
	array := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Array: Before[%s]\n", array[1])

	// value semantic iteration for array
	// shallow copy original array
	for i, v := range array {
		// affect the original array
		// not affect the array in the iteration
		array[1] = "Jack"
		if i == 1 {
			// output: Before[Betty] : In Loop[Betty]
			fmt.Printf("Value Semantic Iteration: In Loop[%s]\n", v)
		}
	}

	// reset array
	array = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	// pointer semantic iteration for array
	// loop the original array
	for i := range array {
		array[1] = "Jack"
		if i == 1 {
			// output: Before[Betty] : In Loop[Jack]
			fmt.Printf("Pointer Semantic Iteration: In Loop[%s]\n", array[i])
		}
	}

	slice := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	// slice is a reference -> loop the original slice
	// but the i is limited to the range of original length
	// and the v binds the original value in this case
	for i, v := range slice {
		if i == 1 {
			slice[i] = "Billy"
		}
		fmt.Println(i, v)
	}
	// reset
	slice = []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for i := range slice {
		if i == 0 {
			slice = append(slice, "Billy")
		}
		fmt.Println(i, slice[i])
		if i == 4 {
			fmt.Println(i+1, slice[i+1])
		}
	}
}
