//记忆化搜索
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	memo := make([]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = -1
	}
	return dfs(cost, n, memo)
}

func dfs(cost []int, i int, memo []int) int {
	if i < 2 {
		return 0
	}
	if memo[i] != -1 {
		return memo[i]
	}
	memo[i] = min(dfs(cost, i-1, memo)+cost[i-1], dfs(cost, i-2, memo)+cost[i-2])
	return memo[i]
}

//递推
func minCostClimbingStairs1(cost []int) int {
	n := len(cost)
	fn := make([]int, n+1)
	for i := 2; i <= n; i++ {
		fn[i] = min(fn[i-1]+cost[i-1], fn[i-2]+cost[i-2])
	}
	return fn[n]
}