package pattern

import "fmt"

func Seven() {
	for i := 1; i < 5+1; i++ {
		for range 5 - i {
			fmt.Printf(" ")
		}
		for range i + (i - 1) {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

}
