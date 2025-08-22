package 矩阵

func rotate(matrix [][]int) {
	m := len(matrix)
	for i := 0; i < m/2; i++ {
		for j := 0; j < (m+1)/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[m-j-1][i]
			matrix[m-j-1][i] = matrix[m-i-1][m-j-1]
			matrix[m-i-1][m-j-1] = matrix[j][m-i-1]
			matrix[j][m-i-1] = tmp
		}
	}
}
