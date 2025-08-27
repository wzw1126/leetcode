func findKthLargest(nums []int, k int) int {
	return sortNums(nums, 0, len(nums)-1, len(nums)-k)
}
func sortNums(nums []int, left, right, x int) int {
	index := parition(nums, left, right)
	if x == index {
		return nums[index]
	} else if index < x {
		return sortNums(nums, index+1, right, x)
	} else {
		return sortNums(nums, left, index-1, x)
	}
}
func parition(nums []int, left, right int) int {
	pivot := nums[left]
	i, j := left, right
	for i < j {
		for i < j && nums[j] > pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}