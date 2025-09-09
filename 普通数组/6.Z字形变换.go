func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	res := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]byte, 0)
	}
	flag := -1
	j := 0
	for i := 0; i < len(s); i++ {
		if j == 0 || j == numRows-1 {
			flag = -flag
		}
		res[j] = append(res[j], s[i])
		j += flag
	}
	ans := make([]byte, 0)
	for i := 0; i < numRows; i++ {
		ans = append(ans, res[i]...)
	}
	return string(ans)
}