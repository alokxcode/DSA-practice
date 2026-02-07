package algorithms

func Bubble_Sort(nums []int) []int {
	length := len(nums)
	for length > 1 {
		for i := 0; i < length-1; i++ {
			j := i + 1
			if nums[i] > nums[j] {
				temp := nums[j]
				nums[j] = nums[i]
				nums[i] = temp
			}
		}
		length--
	}
	return nums
}
