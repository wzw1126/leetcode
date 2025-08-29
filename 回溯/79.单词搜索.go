func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j, idx int) bool
	dfs = func(i, j, idx int) bool {
		if board[i][j] != word[idx] {
			return false
		}
		if idx == len(word)-1 {
			return true
		}
		visited[i][j] = true
		if i+1 < m && !visited[i+1][j] && dfs(i+1, j, idx+1) {
			return true
		}
		if i-1 >= 0 && !visited[i-1][j] && dfs(i-1, j, idx+1) {
			return true
		}
		if j+1 < n && !visited[i][j+1] && dfs(i, j+1, idx+1) {
			return true
		}
		if j-1 >= 0 && !visited[i][j-1] && dfs(i, j-1, idx+1) {
			return true
		}
		visited[i][j] = false
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}