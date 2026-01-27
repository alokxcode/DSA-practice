package pattern

import "fmt"

func Nine() {
	for i := 1; i < 6; i++ {
		for range 5 - i {
			fmt.Printf(" ")
		}
		for range i + (i - 1) {
			fmt.Printf("*")
		}
		fmt.Printf("\n")

	}
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
