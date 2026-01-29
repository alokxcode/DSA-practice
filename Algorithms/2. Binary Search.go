package algorithms

func Binary_Search(nums []int, target int) int {
	// nums should be sorted
	highIndex := len(nums)
	lowIndex := 0

	for lowIndex < highIndex {
		mid := lowIndex + (highIndex-lowIndex)/2
		midV := nums[mid]
		if midV == target {
			return midV
		} else if midV > target {
			highIndex = mid
		} else {
			lowIndex = mid + 1
		}
	}
	return 0

}
