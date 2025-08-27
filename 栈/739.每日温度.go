func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	st := []int{}
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && temperatures[i] >= temperatures[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			res[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	return res
}