func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(s) {
			res = append(res, slices.Clone(path))
			return
		}
		for i := start; i < len(s); i++ {
			if isPalindrome(s, start, i) {
				path = append(path, string(s[start:i+1]))
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}