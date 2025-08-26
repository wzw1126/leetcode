
//使用的开区间二分法
//如果nums[mid] <= nums[n-1]，说明mid在右侧有序数组中
//如果target在右侧有序数组中，则在右侧有序数组中查
//否则在左侧数组中查
func search(nums []int, target int) int {
	n := len(nums)
	left, right := -1, n
	for left+1 < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] <= nums[n-1] {
			if nums[mid] <= target && target <= nums[n-1] {
				left = mid
			} else {
				right = mid
			}
		} else {
			if nums[left+1] <= target && target < nums[mid] {
				right = mid
			} else {
				left = mid
			}
		}
	}
	return -1
}