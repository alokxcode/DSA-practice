package main

import (
	"fmt"

	algorithms "github.com/alokxcode/dsa/Algorithms"
)

func main() {
	sorted_nums := []int{4, 6, 7, 7, 8, 11, 40, 43, 45, 65, 70, 70, 77}
	target := 43

	unsorted_nums := []int{4, 67, 96, 4, 32, 2, 34, 7, 54, 3}

	result := algorithms.Linear_Search(sorted_nums, target)
	fmt.Println(result)

	result2 := algorithms.Binary_Search(sorted_nums, target)
	fmt.Println(result2)

	result3 := algorithms.Bubble_Sort(unsorted_nums)
	fmt.Println(result3)

	result4 := algorithms.Insertion_Sort(unsorted_nums)
	fmt.Println(result4)
}
