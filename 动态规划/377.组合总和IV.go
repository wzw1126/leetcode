func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for s := 1; s <= target; s++ {
		for _, x := range nums {
			if s >= x {
				dp[s] += dp[s-x]
			}
		}
	}
	return dp[target]

}