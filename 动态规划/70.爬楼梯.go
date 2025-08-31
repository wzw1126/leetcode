func climbStairs1(n int) int {
	memo := make([]int, n+1)
	return dfs(memo, n)
}

func dfs(memo []int, n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = dfs(memo, n-1) + dfs(memo, n-2)
	return memo[n]
}

//递推
func climbStairs(n int) int {
	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
