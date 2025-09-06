//二分法
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left := 0
	right := x
	for left+1 < right {
		m := (left + right) / 2
		if m*m <= x {
			left = m
		} else {
			right = m
		}
	}
	return left
}