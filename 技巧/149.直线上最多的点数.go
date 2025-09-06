
/*
假设直线一定经过 points[i]，那么最多还能经过多少个点？
把 points[i]=(x,y) 当作坐标原点（所以直线一定经过坐标原点），其余点 (x2​ ,y2) 的新坐标为 (x2−x,y2−y)。
在直线一定经过坐标原点的情况下，直线斜率为新坐标与原点连线的斜率。斜率相同的点，一定在同一条直线上。
注意斜率不存在的情况（x2−x=0），我们可以把斜率记为一个特殊值（如 math.MaxFloat64）。
遍历所有点，计算斜率并记录斜率出现的次数，取最大值即为经过 points[i] 的直线最多还能经过的点数。
对每个点都执行一次上述操作，取最大值即为最终答案。
*/
func maxPoints(points [][]int) int {
	n := len(points)
	ans := 0
	for i := 0; i < n-1; i++ {
		x1, y1 := points[i][0], points[i][1]
		m := make(map[float64]int, 0)
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			dx := x2 - x1
			dy := y2 - y1
			k := math.MaxFloat64
			if dx != 0 {
				k = float64(dy) / float64(dx)
			}
			m[k]++
			ans = max(ans, m[k])
		}
	}
	return ans + 1
}