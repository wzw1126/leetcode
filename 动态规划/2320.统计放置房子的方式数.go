package main

//单独考虑一侧的房子，定义 f[i] 表示前 i 个地块的放置方案数，其中第 i 个地块可以放房子，也可以不放房子。
//1.不放房子，那么第 i−1 个地块可放可不放，则有 f[i]=f[i−1]。
//2.若放房子，那么第 i−1 个地块无法放房子，第 i−2 个地块可放可不放，则有 f[i]=f[i−2]。
//综上，状态转移方程为 f[i]=f[i−1]+f[i−2]，初始条件为 f[0]=1,f[1]=2。
// (f[0]表示没有地块时只有一种方案，即不放房子，f[1]表示有一个地块时可以选择放或者不放房子，共两种方案)
//最终答案为 f[n] 的平方对 10^9+7 取模的结果，即 (f[n] * f[n]) % (10^9 + 7)。
func countHousePlacements(n int) int {
	memo := make([]int, n+1)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	mod := 1000000007
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 0 {
			return 1
		}
		if i <= 1 {
			return 2
		}
		if memo[i] != -1 {
			return memo[i]
		}
		memo[i] = (dfs(i-1) + dfs(i-2)) % mod
		return memo[i]
	}
	res := dfs(n)
	return (res * res) % mod
}

//递推
func countHousePlacements1(n int) int {
	mod := 1000000007
	fn := make([]int, n+1)
	fn[0] = 1
	fn[1] = 2
	for i := 2; i <= n; i++ {
		fn[i] = (fn[i-1] + fn[i-2]) % mod
	}
	return (fn[n] * fn[n]) % mod
}
