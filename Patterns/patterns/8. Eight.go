package pattern

import "fmt"

func Eight() {
	for i := range 5 {
		for range i {
			fmt.Printf(" ")
		}

		for range 9 - i*2 {
			fmt.Printf("*")
		}

		fmt.Printf("\n")
	}
}
