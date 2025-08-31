func jump(nums []int) int {
	cur := 0
	nextCur := 0
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		nextCur = max(nextCur, i+nums[i])
		if i == cur {
			cur = nextCur
			res++
		}
	}
	return res
}