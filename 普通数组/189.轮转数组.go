package 普通数组

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-k-1)
	reverse(nums, n-k, n-1)
	reverse(nums, 0, n-1)
}

func reverse(nums []int, left int, right int) {
	for left <= right {
		tmp := nums[left]
		nums[left] = nums[right]
		nums[right] = tmp
		left++
		right--
	}
}
