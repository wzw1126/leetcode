//记忆化搜索
func uniquePaths(m int, n int) int {
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if i == 0 && j == 0 {
			return 1
		}
		if memo[i][j] == -1 {
			memo[i][j] = dfs(i-1, j) + dfs(i, j-1)
		}
		return memo[i][j]
	}
	return dfs(m-1, n-1)

}
//递推
func uniquePaths(m int, n int) int {
    dp:=make([][]int,m)
    for i:=0;i<m;i++{
        dp[i]=make([]int,n)
    }
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if i==0 && j==0{
                dp[i][j]=1
                continue
            }
            if i==0 || j==0 {
                if i==0{
                    dp[i][j]=dp[i][j-1]
                }
                if j==0{
                    dp[i][j]=dp[i-1][j]
                }
                continue;
            }
            dp[i][j]=dp[i-1][j]+dp[i][j-1]
            
        }
    }
    return dp[m-1][n-1]
}