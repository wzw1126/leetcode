func rob(nums []int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	return dfs(n-1, nums, memo)
}

func dfs(i int, nums []int, memo []int) int {
	if i < 0 {
		return 0
	}
	if memo[i] != -1 {
		return memo[i]
	}
	node1 := dfs(i-1, nums, memo)
	node2 := dfs(i-2, nums, memo) + nums[i]
	memo[i] = max(node1, node2)
	return memo[i]
}