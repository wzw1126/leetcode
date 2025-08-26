package main

//二分法解法
//由于矩阵的每一行是递增的，且每行的第一个数大于前一行的最后一个数，如果把矩阵每一行拼在一起，我们可以得到一个递增数组。
func searchMatrix1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := -1, m*n
	for left+1 < right {
		mid := left + (right-left)/2
		v := matrix[mid/n][mid%n]
		if v == target {
			return true
		} else if v < target {
			left = mid
		} else {
			right = mid
		}
	}
	return false

}

//排除法
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] < target {
			x++
		} else if matrix[x][y] > target {
			y--
		} else {
			return true
		}
	}
	return false
}
