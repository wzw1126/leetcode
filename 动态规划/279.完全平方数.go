//递推
const N = 10000

var f [101][N + 1]int

func init() {
	for i := 1; i <= N; i++ {
		f[0][i] = math.MaxInt
	}
	for i := 1; i*i <= N; i++ {
		for j := 0; j <= N; j++ {
			if j < i*i {
				f[i][j] = f[i-1][j]
			} else {
				f[i][j] = min(f[i-1][j], f[i][j-i*i]+1)
			}
		}
	}
}
func numSquares(n int) int {
	return f[int(math.Sqrt(float64(n)))][n] // 也可以写 f[100][n]
}

//记忆化搜索
// 写在外面，多个测试数据之间可以共享，减少计算量
var memo [101][10001]int

func init() {
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
}

func dfs(i, j int) int {
	if i == 0 {
		if j == 0 {
			return 0
		}
		return math.MaxInt
	}
	p := &memo[i][j]
	if *p != -1 { // 之前计算过
		return *p
	}
	if j < i*i {
		*p = dfs(i-1, j) // 只能不选
	} else {
		*p = min(dfs(i-1, j), dfs(i, j-i*i)+1) // 不选 vs 选
	}
	return *p
}

func numSquares(n int) int {
	return dfs(int(math.Sqrt(float64(n))), n)
}
