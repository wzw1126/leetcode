func searchRange(nums []int, target int) []int {
	pre := searchIndex(nums, target)
	if pre == len(nums) || nums[pre] != target {
		return []int{-1, -1}
	}
	last := searchIndex(nums, target+1)
	return []int{pre, last - 1}
}

func searchIndex(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}