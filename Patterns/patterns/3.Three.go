package pattern

import "fmt"

func Three() {
	for i := 1; i < 5+1; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%v", j)
		}
		fmt.Printf("\n")
	}
}
