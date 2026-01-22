package pattern

import "fmt"

func Four() {
	for i := 1; i < 5+1; i++ {
		for range i {
			fmt.Printf("%v", i)
		}
		fmt.Printf("\n")
	}
}
