package pattern

import "fmt"

func Two() {
	for i := 1; i < 5+1; i++ {
		for range i {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
