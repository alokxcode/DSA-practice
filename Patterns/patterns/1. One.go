package pattern

import "fmt"

func One() {
	for range 5 {
		for range 5 {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
