func sortColors(nums []int) {
	sortNum(nums, 0, len(nums)-1)
}
func sortNum(nums []int, left, right int) {
	if left >= right {
		return
	}
	index := partion(nums, left, right)
	sortNum(nums, left, index-1)
	sortNum(nums, index+1, right)
}

func partion(nums []int, left, right int) int {
	pivot := nums[left]
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[i] = nums[i], nums[left]
	return i
}

