func removeDuplicates(nums []int) int {
	n := len(nums)
	j := 0
	for i := 1; i < n; i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}