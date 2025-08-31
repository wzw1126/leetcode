func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0], tmp[i] = 1, 1
		for j := 1; j < i; j++ {
			tmp[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = tmp
	}
	return res

}