func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return []string{}
	}
	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	res := make([]string, 0)
	path := make([]byte, 0)
	var dfs func(int)
	dfs = func(start int) {
		if start == len(digits) {
			res = append(res, string(path))
			return
		}
		str := m[digits[start]]
		for i := 0; i < len(str); i++ {
			path = append(path, str[i])
			dfs(start + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

