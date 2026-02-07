package algorithms

func Insertion_Sort(nums []int) []int {
	for i := range nums {
		for j := range i - 1 {
			if i < j {
				temp := nums[j]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}

		// high := i - 1
		// low := 0

		// mid := low + (high-low)/2

		// midV := nums[mid]
		// fmt.Println(midV)
		// for high > low {
		// 	if midV > i {
		// 		high = mid
		// 	} else if midV < i {
		// 		low = mid + 1
		// 	}
		// 	fmt.Println(midV)

		// 	if i == midV {
		// 		temp := nums[midV]
		// 		nums[i] = nums[midV]
		// 		nums[midV] = temp
		// 		break
		// 	}
		// }

	}
	return nums
}
