package hot100

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		tmp := target - v
		if _, ok := m[tmp]; ok {
			return []int{m[tmp], i}
		}
		m[v] = i
	}
	return nil
}
