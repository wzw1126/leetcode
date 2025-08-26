func findMin(nums []int) int {
	n := len(nums)
	left, right := -1, n-1
	for left+1 < right {
		mid := (left + right) / 2
		if nums[mid] < nums[n-1] {
			right = mid
		} else {
			left = mid
		}
	}
	return nums[right]
}