package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		fmt.Printf("*")
	// 	}
	// 	fmt.Printf("\n")
	// }

	// using range
	for range 5 {
		for range 5 {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

	// using recursion

	// printPattern()
}

// func printPattern(x int, y int) string {
// 	if count < x {

// 	}

// }
