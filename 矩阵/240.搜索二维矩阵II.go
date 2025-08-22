package 矩阵

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := n-1, 0
	for x >= 0 && y < m {
		if matrix[y][x] == target {
			return true
		} else if matrix[y][x] > target {
			x--
		} else {
			y++
		}
	}
	return false
}
