package main

import "fmt"

func main() {
	five := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for i := range five {
		if i == 0 {
			five = append(five, "Billy")
			fmt.Println(len(five))
		}
		fmt.Printf("v[%s]\n", five[i+1])
	}
}
