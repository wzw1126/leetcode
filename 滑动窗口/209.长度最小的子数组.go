func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	sum := 0
	j := 0
	ans := n + 1
	for i := 0; i < n; i++ {
		sum += nums[i]
		for sum >= target {
			ans = min(ans, i-j+1)
			sum -= nums[j]
			j++
		}
	}
	if ans == n+1 {
		return 0
	}
	return ans
}