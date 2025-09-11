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

// 先进行转置，再进行每行的翻转
func rotate(matrix [][]int) {
	m := len(matrix)

	// 1. Transpose the matrix
	for i := 0; i < m; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 2. Reverse each row
	for i := 0; i < m; i++ {
		l, r := 0, m-1
		for l < r {
			// Corrected: Swap elements within the i-th row
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			l++
			r--
		}
	}
}
