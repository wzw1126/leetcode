func removeElement(nums []int, val int) int {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}