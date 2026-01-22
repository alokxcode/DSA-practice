package pattern

import "fmt"

// Okay so here, we need to print this pattern

// 12345
// 1234
// 123
// 12
// 1

func Five() {
	for i := 1; i < 5+1; i++ {
		for j := 1; j < (5+1)-i+1; j++ {
			fmt.Printf("%v", j)
		}
		fmt.Printf("\n")
	}

}
