func generateParenthesis(n int) []string {
	res := make([]string, 0)
	path := []byte{}
	cnt1, cnt2 := 0, 0
	var dfs func()
	dfs = func() {
		if cnt2 == n {
			res = append(res, string(path))
			return
		}
		if cnt1 < n {
			path = append(path, '(')
			cnt1++
			dfs()
			cnt1--
			path = path[:len(path)-1]
		}
		if cnt2 < cnt1 {
			path = append(path, ')')
			cnt2++
			dfs()
			cnt2--
			path = path[:len(path)-1]
		}
	}
	dfs()
	return res
}