package hot100

func longestConsecutive(nums []int) int {
	has := map[int]bool{}
	for _, num := range nums {
		has[num] = true
	}
	m := len(has)
	ans := 0
	// 枚举所有数
	for x := range has {
		if _, ok := has[x-1]; ok {
			continue
		}
		cur := x
		cnt := 1
		for has[cur+1] {
			cur++
			cnt++
		}
		ans = max(ans, cnt)
		if ans*2 >= m {
			break
		}
	}
	return ans
}
