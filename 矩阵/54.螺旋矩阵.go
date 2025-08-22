package 矩阵

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := []int{}
	l, r, t, d := 0, n-1, 0, m-1
	for true {
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		if t++; t > d {
			break
		}
		for i := t; i <= d; i++ {
			res = append(res, matrix[i][r])
		}
		if r--; r < l {
			break
		}
		for i := r; i >= l; i-- {
			res = append(res, matrix[d][i])
		}
		if d--; d < t {
			break
		}
		for i := d; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		if l++; l > r {
			break
		}
	}
	return res
}
