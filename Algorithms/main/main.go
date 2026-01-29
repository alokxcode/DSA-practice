package main

import (
	"fmt"

	algorithms "github.com/alokxcode/dsa/Algorithms"
)

func main() {
	nums := []int{4, 6, 7, 7, 8, 11, 40, 43, 45, 65, 70, 70, 77}
	target := 43

	result := algorithms.Linear_Search(nums, target)
	fmt.Println(result)

	result2 := algorithms.Binary_Search(nums, target)
	fmt.Println(result2)
}
