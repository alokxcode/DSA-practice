package pattern

import "fmt"

func Eleven() {
	for i := 1; i < 6; i++ {
		for j := 1; j <= i; j++ {
			if j%2 == 0 {
				fmt.Printf("0")
			} else {
				fmt.Printf("1")
			}
		}
		fmt.Printf("\n")
	}
}
