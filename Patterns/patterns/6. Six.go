package pattern

import "fmt"

func Six() {
	for i := range 5 + 1 {
		for range (5) - i {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

}
