package pattern

import "fmt"

func Ten() {
	for i := 1; i < 6; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	for i := 1; i < 5; i++ {
		for j := 1; j <= 5-i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

}
